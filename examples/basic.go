package main

import (
	"flag"
	"fmt"
	"time"
	"github.com/XiaoMi/galaxy-sdk-go/sds/auth"
	"github.com/XiaoMi/galaxy-sdk-go/sds/common"
	"github.com/XiaoMi/galaxy-sdk-go/sds/client"
	"github.com/XiaoMi/galaxy-sdk-go/sds/table"
	"github.com/XiaoMi/galaxy-sdk-go/thrift"
	"github.com/XiaoMi/galaxy-sdk-go/sds/errors"
	"math/rand"
)

func main() {
	flag.Parse()

	// Set your AppId and AppSecret
	appId := ""
	appSecret := ""
	cred := auth.Credential{auth.UserType_APP_SECRET, appId, thrift.StringPtr(appSecret)}
	endpoint := "http://sds.api.xiaomi.com"

	cfAdmin := client.NewClientFactory(&cred,
		time.Duration(common.DEFAULT_ADMIN_CLIENT_TIMEOUT * int64(time.Second)))
	adminClient := cfAdmin.NewAdminClient(endpoint + common.ADMIN_SERVICE_PATH)
	cfTable := client.NewClientFactory(&cred,
		time.Duration(common.DEFAULT_CLIENT_TIMEOUT * int64(time.Second)))
	tableClient := cfTable.NewTableClient(endpoint + common.TABLE_SERVICE_PATH)

	tableName := "go-test-weather"
	tableSpec := table.TableSpec{
		Schema : &table.TableSchema{
			PrimaryIndex: []*table.KeySpec{&table.KeySpec{"cityId", true}},
			Attributes: map[string]table.DataType {
				"cityId": table.DataType_STRING,
				"timestamp": table.DataType_INT64,
				"score": table.DataType_DOUBLE,
				"pm25": table.DataType_INT64,
			},
			PreSplits: 1, // Must be set with Go SDK, the default 0 is illegal
			Ttl: -1, // Must be set with Go SDK, the default 0 is illegal
		},
		Metadata: &table.TableMetadata{
			Quota: &table.TableQuota{100 * 1024 * 1024}, // 100M
			Throughput: &table.ProvisionThroughput{10, 10},
		},
	}

	// Drop and create table
	fmt.Println("Dropping old table")
	if err := adminClient.DropTable(tableName); err != nil {
		if se, ok := err.(*errors.ServiceException); ok {
			fmt.Printf("Drop table failed, error: %s, details: %s, callId: %s\n",
				se.ErrorCode, *se.Details, *se.CallId)
		}
	}
	fmt.Println("Creating new table")
	_, err := adminClient.CreateTable(tableName, &tableSpec)
	if err != nil {
		se, ok := err.(*errors.ServiceException)
		if ok {
			fmt.Printf("Error: %s, details: %s, callId: %s\n",
				se.ErrorCode, *se.Details, *se.CallId)
		}
	}

	// put data
	cities := [...]string {"北京", "Beihai", "Dalian", "Dandong", "Fuzhou", "Guangzhou",
		"Haikou", "Hankou", "Huangpu", "Jiujiang", "Lianyungang", "Nanjing", "Nantong",
		"Ningbo", "Qingdao", "Qinhuangdao", "Rizhao", "Sanya", "Shanghai", "Shantou",
		"Shenzhen", "Tianjin", "Weihai", "Wenzhou", "Xiamen", "Yangzhou", "Yantai"}

	for i := 0; i < len(cities); i += 1 {
		put := table.PutRequest{
			TableName: tableName,
			Record: map[string]*table.Datum {
				"cityId": client.StringDatum(cities[i]),
				"timestamp": client.Int64Datum(time.Now().Unix()),
				"score": client.Float64Datum(rand.Float64() * 100),
				"pm25": client.Int64Datum(rand.Int63n(100)),
			},
		}
		if pr, err := tableClient.Put(&put); err != nil {
			fmt.Printf("Failed to put record: %s\n", err)
		} else {
			fmt.Printf("Put record: %s, %v\n", cities[i], pr.IsSetSuccess())
		}
	}

	// get data
	i := rand.Intn(len(cities))
	get := table.GetRequest{
		TableName: tableName,
		Keys: map[string]*table.Datum {
			"cityId": client.StringDatum(cities[i]),
		},
	}
	if gr, err := tableClient.Get(&get); err != nil {
		fmt.Printf("Failed to get record: %s\n", err)
	} else if gr.GetItem() != nil {
		fmt.Printf("Get record, city: %s, score: %f\n",
			client.StringValue(gr.GetItem()["cityId"]),
			client.Float64Value(gr.GetItem()["score"]))
	}

	// batch put
	batch := table.NewBatchRequest()
	for i := 0; i < 2; i += 1 {
		put := table.PutRequest{
			TableName: tableName,
			Record: map[string]*table.Datum {
				"cityId": client.StringDatum(cities[i]),
				"timestamp": client.Int64Datum(time.Now().Unix()),
				"score": client.Float64Datum(rand.Float64() * 100),
				"pm25": client.Int64Datum(rand.Int63n(100)),
			},
		}
		batch.Items = append(batch.Items, &table.BatchRequestItem{
				Action: table.BatchOp_PUT,
				Request: &table.Request{
					PutRequest: &put,
				},
			})
	}
	if br, err := tableClient.Batch(batch); err != nil {
		fmt.Printf("Failed to batch put record: %s\n", err)
	} else {
		fmt.Printf("Batch put record: %s\n", br)
	}

	// scan data
	scan := table.ScanRequest{
		TableName: tableName,
		StartKey: nil,
		StopKey: nil,
		Attributes: []string{"cityId", "score"},
		Condition: thrift.StringPtr("score > 0"), // condition to meet
		Limit: 2,
	}

	for item := range client.NewTableScanner(tableClient, &scan).Iter() {
		if datum, err := item.Datum, item.Error; err != nil {
			fmt.Printf("Failed to scan record: %s\n", err)
			break
		} else {
			fmt.Printf("Scanned record, city: %s, score: %f\n",
				client.StringValue(datum["cityId"]),
				client.Float64Value(datum["score"]))
		}
	}

	fmt.Println("Dropping table")
	err = adminClient.DropTable(tableName)
	if err != nil {
		fmt.Printf("Drop table failed: %s\n", err)
	}
}
