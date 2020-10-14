package metric

import (
	"math"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
	"github.com/soniah/gosnmp"
	"github.com/toni-moreno/snmpcollector/pkg/config"
)

const tolerance = .0000000000001

var opt cmp.Option

func init() {
	opt = cmp.Comparer(func(x, y float64) bool {
		diff := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		return (diff / mean) < tolerance
	})
}

//--------------------------------------------------------------
// INTEGER TEST (gosnmp.Integer 0x02) Pdu Value Type = int
//--------------------------------------------------------------

func Test_Integer32_to_FLOAT(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "linux_memTotalSwap",
		FieldName:   "memTotalSwap",
		Description: "",
		BaseOID:     ".1.3.6.1.4.1.2021.4.3.0",
		DataSrcType: "Integer32",
		GetRate:     false,
		Scale:       0.0,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  0, // to Float64
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Integer,
		Value: int(15789),
	}

	met.SetRawData(data, time.Now())

	switch v := met.CookedValue.(type) {
	case float64:
		if v != 15789.0 {
			t.Errorf("Metric error : got [%v] expected [15789.0]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_Integer32_Scale_to_FLOAT(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "linux_memTotalSwap",
		FieldName:   "memTotalSwap",
		Description: "The total amount of swap space configured for this host.",
		BaseOID:     ".1.3.6.1.4.1.2021.4.3.0",
		DataSrcType: "Integer32",
		GetRate:     false,
		Scale:       0.001,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  0, // to Float64
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Integer,
		Value: int(15789),
	}

	met.SetRawData(data, time.Now())

	switch v := met.CookedValue.(type) {
	case float64:
		//15789*0.001 = 15.789
		if v != 15.789 {
			t.Errorf("Metric error : got [%v] expected [15.789]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_Integer32_to_INTEGER(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "linux_memTotalSwap",
		FieldName:   "memTotalSwap",
		Description: "The total amount of swap space configured for this host.",
		BaseOID:     ".1.3.6.1.4.1.2021.4.3.0",
		DataSrcType: "Integer32",
		GetRate:     false,
		Scale:       0.0,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  1, // to integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Integer,
		Value: int(15789), //after scale and Round => 16
	}

	met.SetRawData(data, time.Now())

	switch v := met.CookedValue.(type) {
	case int64:
		if v != 15789 {
			t.Errorf("Metric error : got [%v] expected [16]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_Integer32_Scale_to_INTEGER(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "linux_memTotalSwap",
		FieldName:   "memTotalSwap",
		Description: "The total amount of swap space configured for this host.",
		BaseOID:     ".1.3.6.1.4.1.2021.4.3.0",
		DataSrcType: "Integer32",
		GetRate:     false,
		Scale:       0.001,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  1, // to integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Integer,
		Value: int(15789), //after scale and Round => 16
	}

	met.SetRawData(data, time.Now())

	switch v := met.CookedValue.(type) {
	case int64:
		//15789*0.001 = 15.789 => round(15.789) = 16
		if v != 16 {
			t.Errorf("Metric error : got [%v] expected [16]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

//--------------------------------------------------------------
// Counter32 TEST (gosnmp.Counter32 0x41) Pdu Value Type = uint
//--------------------------------------------------------------

func Test_Counter32_to_FLOAT(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "myCounter32",
		FieldName:   "counter",
		Description: "",
		BaseOID:     ".1.3.6.1.2.1.6.10.0",
		DataSrcType: "Counter32",
		GetRate:     false,
		Scale:       0.0,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  0, // to Float64
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(15800),
	}

	met.SetRawData(data, time.Now())

	switch v := met.CookedValue.(type) {
	case float64:
		if v != 15800.0 {
			t.Errorf("Metric error : got [%v] expected [15800.0]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_Counter32_Scale_to_FLOAT(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "myCounter32",
		FieldName:   "counter",
		Description: "",
		BaseOID:     ".1.3.6.1.2.1.6.10.0",
		DataSrcType: "Counter32",
		GetRate:     false,
		Scale:       0.001,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  0, // to Float64
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(15800),
	}

	met.SetRawData(data, time.Now())

	switch v := met.CookedValue.(type) {
	case float64:
		//15800*0.001 = 15.8
		if v != 15.8 {
			t.Errorf("Metric error : got [%v] expected [15.8]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_Counter32_to_INTEGER(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "myCounter32",
		FieldName:   "counter",
		Description: "",
		BaseOID:     ".1.3.6.1.2.1.6.10.0",
		DataSrcType: "Counter32",
		GetRate:     false,
		Scale:       0.0,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  1, // to integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(15800), //after scale and Round => 16
	}

	met.SetRawData(data, time.Now())

	switch v := met.CookedValue.(type) {
	case int64:
		if v != 15800 {
			t.Errorf("Metric error : got [%v] expected [15800]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_Counter32_Scale_to_INTEGER(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "myCounter32",
		FieldName:   "counter",
		Description: "",
		BaseOID:     ".1.3.6.1.2.1.6.10.0",
		DataSrcType: "Counter32",
		GetRate:     false,
		Scale:       0.001,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  1, // to integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(15800), //after scale and Round => 16
	}

	met.SetRawData(data, time.Now())

	switch v := met.CookedValue.(type) {
	case int64:
		//15800*0.001 = 15.8 => round(15.8) = 16
		if v != 16 {
			t.Errorf("Metric error : got [%v] expected [16]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

//--------------------------------------------------------------
//                COUNTER32 TEST
//			from: gosnmp.Counter32 (0x41) Pdu Value Type = uint
//--------------------------------------------------------------

func Test_COUNTER32_to_FLOAT(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "my_counter32",
		FieldName:   "anycounter",
		Description: "",
		BaseOID:     ".1.3.6.1.2.1.6.10.0",
		DataSrcType: "COUNTER32",
		GetRate:     false,
		Scale:       0.0,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  0, // to float64
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	now := time.Now()
	before := now.Add(-60 * time.Second)
	//1st data
	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(156000),
	}

	met.SetRawData(data, before)

	//2nd data
	data = gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(162600),
	}
	met.SetRawData(data, now)

	switch v := met.CookedValue.(type) {
	case float64:
		//(162600 - 15600)=6600.0
		if v != 6600.0 {
			t.Errorf("Metric error : got [%v] expected [6600.0]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_COUNTER32_to_INTEGER(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "my_counter32",
		FieldName:   "anycounter",
		Description: "",
		BaseOID:     ".1.3.6.1.2.1.6.10.0",
		DataSrcType: "COUNTER32",
		GetRate:     false,
		Scale:       0.0,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  1, // to Integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	now := time.Now()
	before := now.Add(-60 * time.Second)

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(156000),
	}

	met.SetRawData(data, before)

	//2nd data
	data = gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(162600),
	}
	met.SetRawData(data, now)

	switch v := met.CookedValue.(type) {
	case int64:
		//(162600 - 15600)=6600 => round(6600) = 6600
		if v != 6600 {
			t.Errorf("Metric error : got [%v] expected [6600]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_COUNTER32_rate_to_FLOAT(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "my_counter",
		FieldName:   "anycounter",
		Description: "",
		BaseOID:     ".1.3.6.1.2.1.6.10.0",
		DataSrcType: "COUNTER32",
		GetRate:     true,
		Scale:       0.0,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  0, // to Integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	now := time.Now()
	before := now.Add(-60 * time.Second)

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(156000),
	}

	met.SetRawData(data, before)

	//2nd data
	data = gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(162600),
	}
	met.SetRawData(data, now)

	switch v := met.CookedValue.(type) {
	case float64:
		//(162600 - 15600)=6600/60 = 110
		if v != 110.0 {
			t.Errorf("Metric error : got [%f] expected [110.0]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_COUNTER32_Scale_to_FLOAT(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "my_counter",
		FieldName:   "anycounter",
		Description: "The total amount of swap space configured for this host.",
		BaseOID:     ".1.3.6.1.2.1.6.10.0",
		DataSrcType: "COUNTER32",
		GetRate:     false,
		Scale:       0.001,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  0, // to Integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	now := time.Now()
	before := now.Add(-60 * time.Second)

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(156000),
	}

	met.SetRawData(data, before)

	//2nd data
	data = gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(162600),
	}
	met.SetRawData(data, now)

	switch v := met.CookedValue.(type) {
	case float64:
		//(162600 - 15600)=6600*0.001 = 6.600
		if !cmp.Equal(v, 6.6000, opt) {
			//need compare with tolerance the resulting value was 6.6000000000000005
			t.Errorf("Metric error : got [%v] expected [6.6000]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_COUNTER32_Scale_to_INTEGER(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "my_counter",
		FieldName:   "anycounter",
		Description: "The total amount of swap space configured for this host.",
		BaseOID:     ".1.3.6.1.2.1.6.10.0",
		DataSrcType: "COUNTER32",
		GetRate:     false,
		Scale:       0.001,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  1, // to Integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	now := time.Now()
	before := now.Add(-60)

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(156000),
	}

	met.SetRawData(data, before)

	//2nd data
	data = gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.6.10.0",
		Type:  gosnmp.Counter32,
		Value: uint(162600),
	}
	met.SetRawData(data, now)

	switch v := met.CookedValue.(type) {
	case int64:
		//(162600 - 15600)=6600*0.001 = 6.6 => round(6.6) = 7
		if v != 7 {
			t.Errorf("Metric error : got [%v] expected [7]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

//---------------------------------------------------------------
// Counter64 TEST (gosnmp.Counter64 0x46) Pdu Value Type = uint64
//---------------------------------------------------------------

func Test_Counter64_to_FLOAT(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "myCounter64",
		FieldName:   "counter",
		Description: "",
		BaseOID:     ".1.3.6.1.2.1.31.1.1.1.11.8",
		DataSrcType: "Counter64",
		GetRate:     false,
		Scale:       0.0,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  0, // to Float64
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.31.1.1.1.11.8",
		Type:  gosnmp.Counter64,
		Value: uint64(15800),
	}

	met.SetRawData(data, time.Now())

	switch v := met.CookedValue.(type) {
	case float64:
		if v != 15800.0 {
			t.Errorf("Metric error : got [%v] expected [15800.0]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_Counter64_Scale_to_FLOAT(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "myCounter64",
		FieldName:   "counter",
		Description: "",
		BaseOID:     ".1.3.6.1.2.1.31.1.1.1.11.8",
		DataSrcType: "Counter64",
		GetRate:     false,
		Scale:       0.001,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  0, // to Float64
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.31.1.1.1.11.8",
		Type:  gosnmp.Counter64,
		Value: uint64(15800),
	}

	met.SetRawData(data, time.Now())

	switch v := met.CookedValue.(type) {
	case float64:
		//15800*0.001 = 15.8
		if v != 15.8 {
			t.Errorf("Metric error : got [%v] expected [15.8]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_Counter64_to_INTEGER(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "myCounter64",
		FieldName:   "counter",
		Description: "",
		BaseOID:     ".1.3.6.1.2.1.31.1.1.1.11.8",
		DataSrcType: "Counter64",
		GetRate:     false,
		Scale:       0.0,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  1, // to integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.31.1.1.1.11.8",
		Type:  gosnmp.Counter64,
		Value: uint64(15800), //after scale and Round => 16
	}

	met.SetRawData(data, time.Now())

	switch v := met.CookedValue.(type) {
	case int64:
		if v != 15800 {
			t.Errorf("Metric error : got [%v] expected [15800]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_Counter64_Scale_to_INTEGER(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "myCounter64",
		FieldName:   "counter",
		Description: "",
		BaseOID:     ".1.3.6.1.2.1.31.1.1.1.11.8",
		DataSrcType: "Counter64",
		GetRate:     false,
		Scale:       0.001,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  1, // to integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.31.1.1.1.11.8",
		Type:  gosnmp.Counter64,
		Value: uint64(15800), //after scale and Round => 16
	}

	met.SetRawData(data, time.Now())

	switch v := met.CookedValue.(type) {
	case int64:
		//15800*0.001 = 15.8 => round(15.8) = 16
		if v != 16 {
			t.Errorf("Metric error : got [%v] expected [16]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

//--------------------------------------------------------------
//                COUNTER64 TEST
//			from: gosnmp.Counter64 (0x46) Pdu Value Type = uint64
//--------------------------------------------------------------

func Test_COUNTER64_to_FLOAT(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "my_counter",
		FieldName:   "anycounter",
		Description: "The total amount of swap space configured for this host.",
		BaseOID:     ".1.3.6.1.4.1.2021.4.3.0",
		DataSrcType: "COUNTER64",
		GetRate:     false,
		Scale:       0.0,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  0, // to float64
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	now := time.Now()
	before := now.Add(-60 * time.Second)
	//1st data
	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Counter64,
		Value: uint64(156000),
	}

	met.SetRawData(data, before)

	//2nd data
	data = gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Counter64,
		Value: uint64(162600),
	}
	met.SetRawData(data, now)

	switch v := met.CookedValue.(type) {
	case float64:
		if v != 6600.0 {
			t.Errorf("Metric error : got [%v] expected [6600.0]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_COUNTER64_to_INTEGER(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "my_counter",
		FieldName:   "anycounter",
		Description: "The total amount of swap space configured for this host.",
		BaseOID:     ".1.3.6.1.4.1.2021.4.3.0",
		DataSrcType: "COUNTER64",
		GetRate:     false,
		Scale:       0.0,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  1, // to Integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	now := time.Now()
	before := now.Add(-60 * time.Second)
	//1st data
	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Counter64,
		Value: uint64(156000),
	}

	met.SetRawData(data, before)

	//2nd data
	data = gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Counter64,
		Value: uint64(162600),
	}
	met.SetRawData(data, now)

	switch v := met.CookedValue.(type) {
	case int64:
		//(162600 - 15600)=6600=>round(6600)=6600
		if v != 6600 {
			t.Errorf("Metric error : got [%v] expected [6600]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_COUNTER64_rate_to_FLOAT(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "my_counter",
		FieldName:   "anycounter",
		Description: "The total amount of swap space configured for this host.",
		BaseOID:     ".1.3.6.1.4.1.2021.4.3.0",
		DataSrcType: "COUNTER64",
		GetRate:     true,
		Scale:       0.0,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  0, // to Integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	now := time.Now()
	before := now.Add(-60 * time.Second)

	//1st data
	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Counter64,
		Value: uint64(156000),
	}

	met.SetRawData(data, before)

	//2nd data
	data = gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Counter64,
		Value: uint64(162600),
	}
	met.SetRawData(data, now)

	switch v := met.CookedValue.(type) {
	case float64:
		//(162600 - 15600)=6600/60 = 110
		if v != 110.0 {
			t.Errorf("Metric error : got [%f] expected [110.0]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

//https://play.golang.org/p/E_VQv8U7ha
//https://dev.to/juliaferraioli/testing-in-go-testing-floating-point-numbers-4i0a
//https://www.reddit.com/r/golang/comments/5rwywn/float64_precision/
//https://floating-point-gui.de/

func Test_COUNTER64_Scale_to_FLOAT(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "my_counter",
		FieldName:   "anycounter",
		Description: "The total amount of swap space configured for this host.",
		BaseOID:     ".1.3.6.1.4.1.2021.4.3.0",
		DataSrcType: "COUNTER64",
		GetRate:     false,
		Scale:       0.001,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  0, // to Integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	now := time.Now()
	before := now.Add(-60 * time.Second)
	//1st data
	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Counter64,
		Value: uint64(156000),
	}

	met.SetRawData(data, before)

	//2nd data
	data = gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Counter64,
		Value: uint64(162600),
	}
	met.SetRawData(data, now)

	switch v := met.CookedValue.(type) {
	case float64:
		//(162600 - 15600)=6600*0.001 = 6.6
		if !cmp.Equal(v, 6.6000, opt) {
			//need compare with tolerance the resulting value was 6.6000000000000005
			t.Errorf("Metric error : got [%v] expected [6.6000]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}

func Test_COUNTER64_Scale_to_INTEGER(t *testing.T) {
	mc := &config.SnmpMetricCfg{
		ID:          "my_counter",
		FieldName:   "anycounter",
		Description: "The total amount of swap space configured for this host.",
		BaseOID:     ".1.3.6.1.4.1.2021.4.3.0",
		DataSrcType: "COUNTER64",
		GetRate:     false,
		Scale:       0.001,
		Shift:       0.0,
		IsTag:       false,
		ExtraData:   "",
		Conversion:  1, // to Integer
	}

	met := new(SnmpMetric)
	met.Init(mc)
	met.SetLogger(logrus.New())

	now := time.Now()
	before := now.Add(-60)

	//1st data
	data := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Counter64,
		Value: uint64(156000),
	}

	met.SetRawData(data, before)

	//2nd data
	data = gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.4.1.2021.4.3.0",
		Type:  gosnmp.Counter64,
		Value: uint64(162600),
	}
	met.SetRawData(data, now)

	switch v := met.CookedValue.(type) {
	case int64:
		//(162600 - 15600)=6600*0.001 = round(6.6) = 7
		if v != 7 {
			t.Errorf("Metric error : got [%v] expected [7]", v)
			return
		}
		t.Log("OK")
	default:
		t.Errorf("Metric conversion error to [%T] type", v)
	}
}
