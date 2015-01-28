// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package admin

import (
	"bytes"
	"fmt"
	"github.com/XiaoMi/galaxy-sdk-go/sds/auth"
	"github.com/XiaoMi/galaxy-sdk-go/sds/common"
	"github.com/XiaoMi/galaxy-sdk-go/sds/errors"
	"github.com/XiaoMi/galaxy-sdk-go/sds/table"
	"github.com/XiaoMi/galaxy-sdk-go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = errors.GoUnusedProtection__
var _ = common.GoUnusedProtection__
var _ = auth.GoUnusedProtection__
var _ = table.GoUnusedProtection__
var GoUnusedProtection__ int

//系统统计指标类型
type MetricKey int64

const (
	MetricKey_METER_METRIC_MIN     MetricKey = 0
	MetricKey_READ_ALLOWED         MetricKey = 1
	MetricKey_READ_THROTTLED       MetricKey = 2
	MetricKey_WRITE_ALLOWED        MetricKey = 3
	MetricKey_WRITE_THROTTLED      MetricKey = 4
	MetricKey_ACTION_SUCCESS       MetricKey = 5
	MetricKey_ACTION_CLIENT_ERROR  MetricKey = 6
	MetricKey_ACTION_SYSTEM_ERROR  MetricKey = 7
	MetricKey_METER_METRIC_MAX     MetricKey = 49
	MetricKey_HISTOGRAM_METRIC_MIN MetricKey = 50
	MetricKey_CREATE_LATENCY       MetricKey = 51
	MetricKey_DROP_LATENCY         MetricKey = 52
	MetricKey_DESCRIBE_LATENCY     MetricKey = 53
	MetricKey_ALTER_LATENCY        MetricKey = 54
	MetricKey_ENABLE_LATENCY       MetricKey = 55
	MetricKey_DISABLE_LATENCY      MetricKey = 56
	MetricKey_METRICQUERY_LATENCY  MetricKey = 57
	MetricKey_GET_LATENCY          MetricKey = 58
	MetricKey_PUT_LATENCY          MetricKey = 59
	MetricKey_INCREMENT_LATENCY    MetricKey = 60
	MetricKey_DELETE_LATENCY       MetricKey = 61
	MetricKey_SCAN_LATENCY         MetricKey = 62
	MetricKey_BATCH_LATENCY        MetricKey = 63
	MetricKey_HISTOGRAM_METRIC_MAX MetricKey = 100
)

func (p MetricKey) String() string {
	switch p {
	case MetricKey_METER_METRIC_MIN:
		return "MetricKey_METER_METRIC_MIN"
	case MetricKey_READ_ALLOWED:
		return "MetricKey_READ_ALLOWED"
	case MetricKey_READ_THROTTLED:
		return "MetricKey_READ_THROTTLED"
	case MetricKey_WRITE_ALLOWED:
		return "MetricKey_WRITE_ALLOWED"
	case MetricKey_WRITE_THROTTLED:
		return "MetricKey_WRITE_THROTTLED"
	case MetricKey_ACTION_SUCCESS:
		return "MetricKey_ACTION_SUCCESS"
	case MetricKey_ACTION_CLIENT_ERROR:
		return "MetricKey_ACTION_CLIENT_ERROR"
	case MetricKey_ACTION_SYSTEM_ERROR:
		return "MetricKey_ACTION_SYSTEM_ERROR"
	case MetricKey_METER_METRIC_MAX:
		return "MetricKey_METER_METRIC_MAX"
	case MetricKey_HISTOGRAM_METRIC_MIN:
		return "MetricKey_HISTOGRAM_METRIC_MIN"
	case MetricKey_CREATE_LATENCY:
		return "MetricKey_CREATE_LATENCY"
	case MetricKey_DROP_LATENCY:
		return "MetricKey_DROP_LATENCY"
	case MetricKey_DESCRIBE_LATENCY:
		return "MetricKey_DESCRIBE_LATENCY"
	case MetricKey_ALTER_LATENCY:
		return "MetricKey_ALTER_LATENCY"
	case MetricKey_ENABLE_LATENCY:
		return "MetricKey_ENABLE_LATENCY"
	case MetricKey_DISABLE_LATENCY:
		return "MetricKey_DISABLE_LATENCY"
	case MetricKey_METRICQUERY_LATENCY:
		return "MetricKey_METRICQUERY_LATENCY"
	case MetricKey_GET_LATENCY:
		return "MetricKey_GET_LATENCY"
	case MetricKey_PUT_LATENCY:
		return "MetricKey_PUT_LATENCY"
	case MetricKey_INCREMENT_LATENCY:
		return "MetricKey_INCREMENT_LATENCY"
	case MetricKey_DELETE_LATENCY:
		return "MetricKey_DELETE_LATENCY"
	case MetricKey_SCAN_LATENCY:
		return "MetricKey_SCAN_LATENCY"
	case MetricKey_BATCH_LATENCY:
		return "MetricKey_BATCH_LATENCY"
	case MetricKey_HISTOGRAM_METRIC_MAX:
		return "MetricKey_HISTOGRAM_METRIC_MAX"
	}
	return "<UNSET>"
}

func MetricKeyFromString(s string) (MetricKey, error) {
	switch s {
	case "MetricKey_METER_METRIC_MIN":
		return MetricKey_METER_METRIC_MIN, nil
	case "MetricKey_READ_ALLOWED":
		return MetricKey_READ_ALLOWED, nil
	case "MetricKey_READ_THROTTLED":
		return MetricKey_READ_THROTTLED, nil
	case "MetricKey_WRITE_ALLOWED":
		return MetricKey_WRITE_ALLOWED, nil
	case "MetricKey_WRITE_THROTTLED":
		return MetricKey_WRITE_THROTTLED, nil
	case "MetricKey_ACTION_SUCCESS":
		return MetricKey_ACTION_SUCCESS, nil
	case "MetricKey_ACTION_CLIENT_ERROR":
		return MetricKey_ACTION_CLIENT_ERROR, nil
	case "MetricKey_ACTION_SYSTEM_ERROR":
		return MetricKey_ACTION_SYSTEM_ERROR, nil
	case "MetricKey_METER_METRIC_MAX":
		return MetricKey_METER_METRIC_MAX, nil
	case "MetricKey_HISTOGRAM_METRIC_MIN":
		return MetricKey_HISTOGRAM_METRIC_MIN, nil
	case "MetricKey_CREATE_LATENCY":
		return MetricKey_CREATE_LATENCY, nil
	case "MetricKey_DROP_LATENCY":
		return MetricKey_DROP_LATENCY, nil
	case "MetricKey_DESCRIBE_LATENCY":
		return MetricKey_DESCRIBE_LATENCY, nil
	case "MetricKey_ALTER_LATENCY":
		return MetricKey_ALTER_LATENCY, nil
	case "MetricKey_ENABLE_LATENCY":
		return MetricKey_ENABLE_LATENCY, nil
	case "MetricKey_DISABLE_LATENCY":
		return MetricKey_DISABLE_LATENCY, nil
	case "MetricKey_METRICQUERY_LATENCY":
		return MetricKey_METRICQUERY_LATENCY, nil
	case "MetricKey_GET_LATENCY":
		return MetricKey_GET_LATENCY, nil
	case "MetricKey_PUT_LATENCY":
		return MetricKey_PUT_LATENCY, nil
	case "MetricKey_INCREMENT_LATENCY":
		return MetricKey_INCREMENT_LATENCY, nil
	case "MetricKey_DELETE_LATENCY":
		return MetricKey_DELETE_LATENCY, nil
	case "MetricKey_SCAN_LATENCY":
		return MetricKey_SCAN_LATENCY, nil
	case "MetricKey_BATCH_LATENCY":
		return MetricKey_BATCH_LATENCY, nil
	case "MetricKey_HISTOGRAM_METRIC_MAX":
		return MetricKey_HISTOGRAM_METRIC_MAX, nil
	}
	return MetricKey(0), fmt.Errorf("not a valid MetricKey string")
}

func MetricKeyPtr(v MetricKey) *MetricKey { return &v }

//统计指标的子类型
//(MetricKey, MetricType) 元组唯一确定一个统计指标
type MetricType int64

const (
	MetricType_COUNT    MetricType = 1
	MetricType_M1_RATE  MetricType = 2
	MetricType_M5_RATE  MetricType = 3
	MetricType_M15_RATE MetricType = 4
	MetricType_MEAN     MetricType = 5
	MetricType_STDDEV   MetricType = 6
	MetricType_P50      MetricType = 7
	MetricType_P75      MetricType = 8
	MetricType_P95      MetricType = 9
	MetricType_P98      MetricType = 10
	MetricType_P99      MetricType = 11
)

func (p MetricType) String() string {
	switch p {
	case MetricType_COUNT:
		return "MetricType_COUNT"
	case MetricType_M1_RATE:
		return "MetricType_M1_RATE"
	case MetricType_M5_RATE:
		return "MetricType_M5_RATE"
	case MetricType_M15_RATE:
		return "MetricType_M15_RATE"
	case MetricType_MEAN:
		return "MetricType_MEAN"
	case MetricType_STDDEV:
		return "MetricType_STDDEV"
	case MetricType_P50:
		return "MetricType_P50"
	case MetricType_P75:
		return "MetricType_P75"
	case MetricType_P95:
		return "MetricType_P95"
	case MetricType_P98:
		return "MetricType_P98"
	case MetricType_P99:
		return "MetricType_P99"
	}
	return "<UNSET>"
}

func MetricTypeFromString(s string) (MetricType, error) {
	switch s {
	case "MetricType_COUNT":
		return MetricType_COUNT, nil
	case "MetricType_M1_RATE":
		return MetricType_M1_RATE, nil
	case "MetricType_M5_RATE":
		return MetricType_M5_RATE, nil
	case "MetricType_M15_RATE":
		return MetricType_M15_RATE, nil
	case "MetricType_MEAN":
		return MetricType_MEAN, nil
	case "MetricType_STDDEV":
		return MetricType_STDDEV, nil
	case "MetricType_P50":
		return MetricType_P50, nil
	case "MetricType_P75":
		return MetricType_P75, nil
	case "MetricType_P95":
		return MetricType_P95, nil
	case "MetricType_P98":
		return MetricType_P98, nil
	case "MetricType_P99":
		return MetricType_P99, nil
	}
	return MetricType(0), fmt.Errorf("not a valid MetricType string")
}

func MetricTypePtr(v MetricType) *MetricType { return &v }

//时间间隔单位，用于查询统计指标时的下采样
type TimeSpanUnit int64

const (
	TimeSpanUnit_SECONDS TimeSpanUnit = 1
	TimeSpanUnit_MINUTES TimeSpanUnit = 2
	TimeSpanUnit_HOURS   TimeSpanUnit = 3
)

func (p TimeSpanUnit) String() string {
	switch p {
	case TimeSpanUnit_SECONDS:
		return "TimeSpanUnit_SECONDS"
	case TimeSpanUnit_MINUTES:
		return "TimeSpanUnit_MINUTES"
	case TimeSpanUnit_HOURS:
		return "TimeSpanUnit_HOURS"
	}
	return "<UNSET>"
}

func TimeSpanUnitFromString(s string) (TimeSpanUnit, error) {
	switch s {
	case "TimeSpanUnit_SECONDS":
		return TimeSpanUnit_SECONDS, nil
	case "TimeSpanUnit_MINUTES":
		return TimeSpanUnit_MINUTES, nil
	case "TimeSpanUnit_HOURS":
		return TimeSpanUnit_HOURS, nil
	}
	return TimeSpanUnit(0), fmt.Errorf("not a valid TimeSpanUnit string")
}

func TimeSpanUnitPtr(v TimeSpanUnit) *TimeSpanUnit { return &v }

type AppInfo struct {
	AppId           *string           `thrift:"appId,1" json:"appId"`
	DeveloperId     *string           `thrift:"developerId,2" json:"developerId"`
	TableMappings   map[string]string `thrift:"tableMappings,3" json:"tableMappings"`
	OauthAppMapping map[string]string `thrift:"oauthAppMapping,4" json:"oauthAppMapping"`
	AppName         *string           `thrift:"appName,5" json:"appName"`
}

func NewAppInfo() *AppInfo {
	return &AppInfo{}
}

var AppInfo_AppId_DEFAULT string

func (p *AppInfo) GetAppId() string {
	if !p.IsSetAppId() {
		return AppInfo_AppId_DEFAULT
	}
	return *p.AppId
}

var AppInfo_DeveloperId_DEFAULT string

func (p *AppInfo) GetDeveloperId() string {
	if !p.IsSetDeveloperId() {
		return AppInfo_DeveloperId_DEFAULT
	}
	return *p.DeveloperId
}

var AppInfo_TableMappings_DEFAULT map[string]string

func (p *AppInfo) GetTableMappings() map[string]string {
	return p.TableMappings
}

var AppInfo_OauthAppMapping_DEFAULT map[string]string

func (p *AppInfo) GetOauthAppMapping() map[string]string {
	return p.OauthAppMapping
}

var AppInfo_AppName_DEFAULT string

func (p *AppInfo) GetAppName() string {
	if !p.IsSetAppName() {
		return AppInfo_AppName_DEFAULT
	}
	return *p.AppName
}
func (p *AppInfo) IsSetAppId() bool {
	return p.AppId != nil
}

func (p *AppInfo) IsSetDeveloperId() bool {
	return p.DeveloperId != nil
}

func (p *AppInfo) IsSetTableMappings() bool {
	return p.TableMappings != nil
}

func (p *AppInfo) IsSetOauthAppMapping() bool {
	return p.OauthAppMapping != nil
}

func (p *AppInfo) IsSetAppName() bool {
	return p.AppName != nil
}

func (p *AppInfo) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AppInfo) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.AppId = &v
	}
	return nil
}

func (p *AppInfo) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.DeveloperId = &v
	}
	return nil
}

func (p *AppInfo) ReadField3(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s", err)
	}
	tMap := make(map[string]string, size)
	p.TableMappings = tMap
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_key0 = v
		}
		var _val1 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_val1 = v
		}
		p.TableMappings[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s", err)
	}
	return nil
}

func (p *AppInfo) ReadField4(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s", err)
	}
	tMap := make(map[string]string, size)
	p.OauthAppMapping = tMap
	for i := 0; i < size; i++ {
		var _key2 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_key2 = v
		}
		var _val3 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_val3 = v
		}
		p.OauthAppMapping[_key2] = _val3
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s", err)
	}
	return nil
}

func (p *AppInfo) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 5: %s", err)
	} else {
		p.AppName = &v
	}
	return nil
}

func (p *AppInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("AppInfo"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *AppInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetAppId() {
		if err := oprot.WriteFieldBegin("appId", thrift.STRING, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:appId: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.AppId)); err != nil {
			return fmt.Errorf("%T.appId (1) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:appId: %s", p, err)
		}
	}
	return err
}

func (p *AppInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetDeveloperId() {
		if err := oprot.WriteFieldBegin("developerId", thrift.STRING, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:developerId: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.DeveloperId)); err != nil {
			return fmt.Errorf("%T.developerId (2) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:developerId: %s", p, err)
		}
	}
	return err
}

func (p *AppInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetTableMappings() {
		if err := oprot.WriteFieldBegin("tableMappings", thrift.MAP, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:tableMappings: %s", p, err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.TableMappings)); err != nil {
			return fmt.Errorf("error writing map begin: %s", err)
		}
		for k, v := range p.TableMappings {
			if err := oprot.WriteString(string(k)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
			if err := oprot.WriteString(string(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return fmt.Errorf("error writing map end: %s", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:tableMappings: %s", p, err)
		}
	}
	return err
}

func (p *AppInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetOauthAppMapping() {
		if err := oprot.WriteFieldBegin("oauthAppMapping", thrift.MAP, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:oauthAppMapping: %s", p, err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.OauthAppMapping)); err != nil {
			return fmt.Errorf("error writing map begin: %s", err)
		}
		for k, v := range p.OauthAppMapping {
			if err := oprot.WriteString(string(k)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
			if err := oprot.WriteString(string(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return fmt.Errorf("error writing map end: %s", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:oauthAppMapping: %s", p, err)
		}
	}
	return err
}

func (p *AppInfo) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetAppName() {
		if err := oprot.WriteFieldBegin("appName", thrift.STRING, 5); err != nil {
			return fmt.Errorf("%T write field begin error 5:appName: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.AppName)); err != nil {
			return fmt.Errorf("%T.appName (5) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 5:appName: %s", p, err)
		}
	}
	return err
}

func (p *AppInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AppInfo(%+v)", *p)
}

type MetricQueryRequest struct {
	TableName          *string       `thrift:"tableName,1" json:"tableName"`
	StartTime          *int64        `thrift:"startTime,2" json:"startTime"`
	StopTime           *int64        `thrift:"stopTime,3" json:"stopTime"`
	MetricKey          *MetricKey    `thrift:"metricKey,4" json:"metricKey"`
	MetricType         *MetricType   `thrift:"metricType,5" json:"metricType"`
	DownsampleInterval *int32        `thrift:"downsampleInterval,6" json:"downsampleInterval"`
	DownsampleTimeUnit *TimeSpanUnit `thrift:"downsampleTimeUnit,7" json:"downsampleTimeUnit"`
}

func NewMetricQueryRequest() *MetricQueryRequest {
	return &MetricQueryRequest{}
}

var MetricQueryRequest_TableName_DEFAULT string

func (p *MetricQueryRequest) GetTableName() string {
	if !p.IsSetTableName() {
		return MetricQueryRequest_TableName_DEFAULT
	}
	return *p.TableName
}

var MetricQueryRequest_StartTime_DEFAULT int64

func (p *MetricQueryRequest) GetStartTime() int64 {
	if !p.IsSetStartTime() {
		return MetricQueryRequest_StartTime_DEFAULT
	}
	return *p.StartTime
}

var MetricQueryRequest_StopTime_DEFAULT int64

func (p *MetricQueryRequest) GetStopTime() int64 {
	if !p.IsSetStopTime() {
		return MetricQueryRequest_StopTime_DEFAULT
	}
	return *p.StopTime
}

var MetricQueryRequest_MetricKey_DEFAULT MetricKey

func (p *MetricQueryRequest) GetMetricKey() MetricKey {
	if !p.IsSetMetricKey() {
		return MetricQueryRequest_MetricKey_DEFAULT
	}
	return *p.MetricKey
}

var MetricQueryRequest_MetricType_DEFAULT MetricType

func (p *MetricQueryRequest) GetMetricType() MetricType {
	if !p.IsSetMetricType() {
		return MetricQueryRequest_MetricType_DEFAULT
	}
	return *p.MetricType
}

var MetricQueryRequest_DownsampleInterval_DEFAULT int32

func (p *MetricQueryRequest) GetDownsampleInterval() int32 {
	if !p.IsSetDownsampleInterval() {
		return MetricQueryRequest_DownsampleInterval_DEFAULT
	}
	return *p.DownsampleInterval
}

var MetricQueryRequest_DownsampleTimeUnit_DEFAULT TimeSpanUnit

func (p *MetricQueryRequest) GetDownsampleTimeUnit() TimeSpanUnit {
	if !p.IsSetDownsampleTimeUnit() {
		return MetricQueryRequest_DownsampleTimeUnit_DEFAULT
	}
	return *p.DownsampleTimeUnit
}
func (p *MetricQueryRequest) IsSetTableName() bool {
	return p.TableName != nil
}

func (p *MetricQueryRequest) IsSetStartTime() bool {
	return p.StartTime != nil
}

func (p *MetricQueryRequest) IsSetStopTime() bool {
	return p.StopTime != nil
}

func (p *MetricQueryRequest) IsSetMetricKey() bool {
	return p.MetricKey != nil
}

func (p *MetricQueryRequest) IsSetMetricType() bool {
	return p.MetricType != nil
}

func (p *MetricQueryRequest) IsSetDownsampleInterval() bool {
	return p.DownsampleInterval != nil
}

func (p *MetricQueryRequest) IsSetDownsampleTimeUnit() bool {
	return p.DownsampleTimeUnit != nil
}

func (p *MetricQueryRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.ReadField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.ReadField7(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *MetricQueryRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.TableName = &v
	}
	return nil
}

func (p *MetricQueryRequest) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.StartTime = &v
	}
	return nil
}

func (p *MetricQueryRequest) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.StopTime = &v
	}
	return nil
}

func (p *MetricQueryRequest) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		temp := MetricKey(v)
		p.MetricKey = &temp
	}
	return nil
}

func (p *MetricQueryRequest) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 5: %s", err)
	} else {
		temp := MetricType(v)
		p.MetricType = &temp
	}
	return nil
}

func (p *MetricQueryRequest) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 6: %s", err)
	} else {
		p.DownsampleInterval = &v
	}
	return nil
}

func (p *MetricQueryRequest) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 7: %s", err)
	} else {
		temp := TimeSpanUnit(v)
		p.DownsampleTimeUnit = &temp
	}
	return nil
}

func (p *MetricQueryRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("MetricQueryRequest"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *MetricQueryRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetTableName() {
		if err := oprot.WriteFieldBegin("tableName", thrift.STRING, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:tableName: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.TableName)); err != nil {
			return fmt.Errorf("%T.tableName (1) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:tableName: %s", p, err)
		}
	}
	return err
}

func (p *MetricQueryRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetStartTime() {
		if err := oprot.WriteFieldBegin("startTime", thrift.I64, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:startTime: %s", p, err)
		}
		if err := oprot.WriteI64(int64(*p.StartTime)); err != nil {
			return fmt.Errorf("%T.startTime (2) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:startTime: %s", p, err)
		}
	}
	return err
}

func (p *MetricQueryRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetStopTime() {
		if err := oprot.WriteFieldBegin("stopTime", thrift.I64, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:stopTime: %s", p, err)
		}
		if err := oprot.WriteI64(int64(*p.StopTime)); err != nil {
			return fmt.Errorf("%T.stopTime (3) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:stopTime: %s", p, err)
		}
	}
	return err
}

func (p *MetricQueryRequest) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetMetricKey() {
		if err := oprot.WriteFieldBegin("metricKey", thrift.I32, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:metricKey: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.MetricKey)); err != nil {
			return fmt.Errorf("%T.metricKey (4) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:metricKey: %s", p, err)
		}
	}
	return err
}

func (p *MetricQueryRequest) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetMetricType() {
		if err := oprot.WriteFieldBegin("metricType", thrift.I32, 5); err != nil {
			return fmt.Errorf("%T write field begin error 5:metricType: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.MetricType)); err != nil {
			return fmt.Errorf("%T.metricType (5) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 5:metricType: %s", p, err)
		}
	}
	return err
}

func (p *MetricQueryRequest) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetDownsampleInterval() {
		if err := oprot.WriteFieldBegin("downsampleInterval", thrift.I32, 6); err != nil {
			return fmt.Errorf("%T write field begin error 6:downsampleInterval: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.DownsampleInterval)); err != nil {
			return fmt.Errorf("%T.downsampleInterval (6) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 6:downsampleInterval: %s", p, err)
		}
	}
	return err
}

func (p *MetricQueryRequest) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetDownsampleTimeUnit() {
		if err := oprot.WriteFieldBegin("downsampleTimeUnit", thrift.I32, 7); err != nil {
			return fmt.Errorf("%T write field begin error 7:downsampleTimeUnit: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.DownsampleTimeUnit)); err != nil {
			return fmt.Errorf("%T.downsampleTimeUnit (7) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 7:downsampleTimeUnit: %s", p, err)
		}
	}
	return err
}

func (p *MetricQueryRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MetricQueryRequest(%+v)", *p)
}

type TimeSeriesData struct {
	TableName  *string           `thrift:"tableName,1" json:"tableName"`
	MetricKey  *MetricKey        `thrift:"metricKey,2" json:"metricKey"`
	MetricType *MetricType       `thrift:"metricType,3" json:"metricType"`
	Data       map[int64]float64 `thrift:"data,4" json:"data"`
}

func NewTimeSeriesData() *TimeSeriesData {
	return &TimeSeriesData{}
}

var TimeSeriesData_TableName_DEFAULT string

func (p *TimeSeriesData) GetTableName() string {
	if !p.IsSetTableName() {
		return TimeSeriesData_TableName_DEFAULT
	}
	return *p.TableName
}

var TimeSeriesData_MetricKey_DEFAULT MetricKey

func (p *TimeSeriesData) GetMetricKey() MetricKey {
	if !p.IsSetMetricKey() {
		return TimeSeriesData_MetricKey_DEFAULT
	}
	return *p.MetricKey
}

var TimeSeriesData_MetricType_DEFAULT MetricType

func (p *TimeSeriesData) GetMetricType() MetricType {
	if !p.IsSetMetricType() {
		return TimeSeriesData_MetricType_DEFAULT
	}
	return *p.MetricType
}

var TimeSeriesData_Data_DEFAULT map[int64]float64

func (p *TimeSeriesData) GetData() map[int64]float64 {
	return p.Data
}
func (p *TimeSeriesData) IsSetTableName() bool {
	return p.TableName != nil
}

func (p *TimeSeriesData) IsSetMetricKey() bool {
	return p.MetricKey != nil
}

func (p *TimeSeriesData) IsSetMetricType() bool {
	return p.MetricType != nil
}

func (p *TimeSeriesData) IsSetData() bool {
	return p.Data != nil
}

func (p *TimeSeriesData) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *TimeSeriesData) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.TableName = &v
	}
	return nil
}

func (p *TimeSeriesData) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		temp := MetricKey(v)
		p.MetricKey = &temp
	}
	return nil
}

func (p *TimeSeriesData) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		temp := MetricType(v)
		p.MetricType = &temp
	}
	return nil
}

func (p *TimeSeriesData) ReadField4(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s", err)
	}
	tMap := make(map[int64]float64, size)
	p.Data = tMap
	for i := 0; i < size; i++ {
		var _key4 int64
		if v, err := iprot.ReadI64(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_key4 = v
		}
		var _val5 float64
		if v, err := iprot.ReadDouble(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_val5 = v
		}
		p.Data[_key4] = _val5
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s", err)
	}
	return nil
}

func (p *TimeSeriesData) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TimeSeriesData"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *TimeSeriesData) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetTableName() {
		if err := oprot.WriteFieldBegin("tableName", thrift.STRING, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:tableName: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.TableName)); err != nil {
			return fmt.Errorf("%T.tableName (1) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:tableName: %s", p, err)
		}
	}
	return err
}

func (p *TimeSeriesData) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetMetricKey() {
		if err := oprot.WriteFieldBegin("metricKey", thrift.I32, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:metricKey: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.MetricKey)); err != nil {
			return fmt.Errorf("%T.metricKey (2) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:metricKey: %s", p, err)
		}
	}
	return err
}

func (p *TimeSeriesData) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetMetricType() {
		if err := oprot.WriteFieldBegin("metricType", thrift.I32, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:metricType: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.MetricType)); err != nil {
			return fmt.Errorf("%T.metricType (3) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:metricType: %s", p, err)
		}
	}
	return err
}

func (p *TimeSeriesData) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetData() {
		if err := oprot.WriteFieldBegin("data", thrift.MAP, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:data: %s", p, err)
		}
		if err := oprot.WriteMapBegin(thrift.I64, thrift.DOUBLE, len(p.Data)); err != nil {
			return fmt.Errorf("error writing map begin: %s", err)
		}
		for k, v := range p.Data {
			if err := oprot.WriteI64(int64(k)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
			if err := oprot.WriteDouble(float64(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return fmt.Errorf("error writing map end: %s", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:data: %s", p, err)
		}
	}
	return err
}

func (p *TimeSeriesData) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TimeSeriesData(%+v)", *p)
}
