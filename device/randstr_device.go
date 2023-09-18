package device

import (
	"encoding/json"
	"k8s.io/klog/v2"
	"math/rand"
	"time"
)

type RandStrDevice struct {
}

type RandStrDeviceProtocolCommon struct {
	CustomizedValues struct {
		ProtocolID int64 `json:"protocolID"`
	} `json:"customizedValues"`
}

type RandStrDeviceConfigData struct {
	Length int64 `json:"length"`
}

type RandStrDeviceVisitor struct {
	ProtocolName string                  `json:"protocolName"`
	ConfigData   RandStrDeviceConfigData `json:"configData"`
}

type RandStrDeviceProtocol struct {
	ProtocolName string                  `json:"protocolName"`
	ConfigData   RandStrDeviceConfigData `json:"configData"`
}

// InitDevice provide configmap parsing to specific protocols
func (d *RandStrDevice) InitDevice(protocolCommon []byte) (err error) {
	klog.V(2).Info("RandStrDevice InitDevice Called with protocolCommon: ", string(protocolCommon))
	protocolCommonObj := RandStrDeviceProtocolCommon{}
	json.Unmarshal(protocolCommon, &protocolCommonObj)
	klog.V(2).Info("RandStrDevice InitDevice Called with protocolCommon", protocolCommon, " json unmarshal success")
	return nil
}

// ReadDeviceData  is an interface that reads data from a specific device, data is a type of string
func (d *RandStrDevice) ReadDeviceData(protocolCommon []byte, visitor []byte, protocol []byte) (data interface{}, err error) {
	klog.V(2).Info("RandStrDevice ReadDeviceData Called with protocolCommon: ", string(protocolCommon), ", visitor: ", string(visitor), ", protocol: ", string(protocol))

	protocolCommonObj := RandStrDeviceProtocolCommon{}
	visitorObj := RandStrDeviceVisitor{}
	protocolObj := RandStrDeviceProtocol{}

	json.Unmarshal(protocolCommon, &protocolCommonObj)
	json.Unmarshal(visitor, &visitorObj)
	json.Unmarshal(protocol, &protocolObj)

	klog.V(2).Info("RandStrDevice ReadDeviceData Called with protocolCommon", string(protocolCommon), ", visitor: ", string(visitor), ", protocol: ", string(protocol), " json unmarshal success")
	strLen := protocolObj.ConfigData.Length
	randStr := d.generateRandStr(strLen)

	klog.V(2).Info("RandStrDevice ReadDeviceData Called with randStr: ", randStr)

	return randStr, nil
}

// WriteDeviceData is an interface that write data to a specific device, data's DataType is Consistent with configmap
func (d *RandStrDevice) WriteDeviceData(data interface{}, protocolCommon []byte, visitor []byte, protocol []byte) (err error) {
	klog.V(2).Info("RandStrDevice WriteDeviceData Called with protocolCommon: ", string(protocolCommon), ", visitor: ", string(visitor), ", protocol: ", string(protocol))
	klog.V(2).Info("Writing is not supported for this device")
	return nil
}

// StopDevice is an interface to stop all devices
func (d *RandStrDevice) StopDevice() (err error) {
	klog.V(2).Info("RandStrDevice StopDevice Called")
	return nil
}

// GetDeviceStatus is an interface to get the device status true is OK , false is DISCONNECTED
func (d *RandStrDevice) GetDeviceStatus(protocolCommon []byte, visitor []byte, protocol []byte) (status bool) {
	klog.V(2).Info("RandStrDevice GetDeviceStatus Called")
	return true
}

func (d *RandStrDevice) generateRandStr(l int64) (str string) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano()) // initialize global pseudo random generator
	result := ""
	for i := 0; i < int(l); i++ {
		idx := rand.Intn(len(charset))
		result += string(charset[idx])
	}
	return result
}
