package main

//#include "moc.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QmlBridge_ITF interface {
	core.QObject_ITF
	QmlBridge_PTR() *QmlBridge
}

func PointerFromQmlBridge(ptr QmlBridge_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlBridge_PTR().Pointer()
	}
	return nil
}

func NewQmlBridgeFromPointer(ptr unsafe.Pointer) *QmlBridge {
	var n = new(QmlBridge)
	n.SetPointer(ptr)
	return n
}

func newQmlBridgeFromPointer(ptr unsafe.Pointer) *QmlBridge {
	var n = NewQmlBridgeFromPointer(ptr)
	for len(n.ObjectName()) < len("QmlBridge_") {
		n.SetObjectName("QmlBridge_" + qt.Identifier())
	}
	return n
}

func (ptr *QmlBridge) QmlBridge_PTR() *QmlBridge {
	return ptr
}

func (ptr *QmlBridge) ConnectSendToQml(f func(source string, action string, data string)) {
	defer qt.Recovering("connect QmlBridge::sendToQml")

	if ptr.Pointer() != nil {
		C.QmlBridge_ConnectSendToQml(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sendToQml", f)
	}
}

func (ptr *QmlBridge) DisconnectSendToQml() {
	defer qt.Recovering("disconnect QmlBridge::sendToQml")

	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectSendToQml(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sendToQml")
	}
}

//export callbackQmlBridgeSendToQml
func callbackQmlBridgeSendToQml(ptr unsafe.Pointer, ptrName *C.char, source *C.char, action *C.char, data *C.char) {
	defer qt.Recovering("callback QmlBridge::sendToQml")

	if signal := qt.GetSignal(C.GoString(ptrName), "sendToQml"); signal != nil {
		signal.(func(string, string, string))(C.GoString(source), C.GoString(action), C.GoString(data))
	}

}

func (ptr *QmlBridge) SendToQml(source string, action string, data string) {
	defer qt.Recovering("QmlBridge::sendToQml")

	if ptr.Pointer() != nil {
		C.QmlBridge_SendToQml(ptr.Pointer(), C.CString(source), C.CString(action), C.CString(data))
	}
}

func (ptr *QmlBridge) ConnectSendToGo(f func(source string, action string, data string)) {
	defer qt.Recovering("connect QmlBridge::sendToGo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sendToGo", f)
	}
}

func (ptr *QmlBridge) DisconnectSendToGo() {
	defer qt.Recovering("disconnect QmlBridge::sendToGo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sendToGo")
	}
}

//export callbackQmlBridgeSendToGo
func callbackQmlBridgeSendToGo(ptr unsafe.Pointer, ptrName *C.char, source *C.char, action *C.char, data *C.char) {
	defer qt.Recovering("callback QmlBridge::sendToGo")

	if signal := qt.GetSignal(C.GoString(ptrName), "sendToGo"); signal != nil {
		signal.(func(string, string, string))(C.GoString(source), C.GoString(action), C.GoString(data))
	}

}

func (ptr *QmlBridge) SendToGo(source string, action string, data string) {
	defer qt.Recovering("QmlBridge::sendToGo")

	if ptr.Pointer() != nil {
		C.QmlBridge_SendToGo(ptr.Pointer(), C.CString(source), C.CString(action), C.CString(data))
	}
}

func (ptr *QmlBridge) ConnectRegisterToGo(f func(object *core.QObject)) {
	defer qt.Recovering("connect QmlBridge::registerToGo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "registerToGo", f)
	}
}

func (ptr *QmlBridge) DisconnectRegisterToGo() {
	defer qt.Recovering("disconnect QmlBridge::registerToGo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "registerToGo")
	}
}

//export callbackQmlBridgeRegisterToGo
func callbackQmlBridgeRegisterToGo(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::registerToGo")

	if signal := qt.GetSignal(C.GoString(ptrName), "registerToGo"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	}

}

func (ptr *QmlBridge) RegisterToGo(object core.QObject_ITF) {
	defer qt.Recovering("QmlBridge::registerToGo")

	if ptr.Pointer() != nil {
		C.QmlBridge_RegisterToGo(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QmlBridge) ConnectDeregisterToGo(f func(objectName string)) {
	defer qt.Recovering("connect QmlBridge::deregisterToGo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deregisterToGo", f)
	}
}

func (ptr *QmlBridge) DisconnectDeregisterToGo() {
	defer qt.Recovering("disconnect QmlBridge::deregisterToGo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deregisterToGo")
	}
}

//export callbackQmlBridgeDeregisterToGo
func callbackQmlBridgeDeregisterToGo(ptr unsafe.Pointer, ptrName *C.char, objectName *C.char) {
	defer qt.Recovering("callback QmlBridge::deregisterToGo")

	if signal := qt.GetSignal(C.GoString(ptrName), "deregisterToGo"); signal != nil {
		signal.(func(string))(C.GoString(objectName))
	}

}

func (ptr *QmlBridge) DeregisterToGo(objectName string) {
	defer qt.Recovering("QmlBridge::deregisterToGo")

	if ptr.Pointer() != nil {
		C.QmlBridge_DeregisterToGo(ptr.Pointer(), C.CString(objectName))
	}
}

func NewQmlBridge(parent core.QObject_ITF) *QmlBridge {
	defer qt.Recovering("QmlBridge::QmlBridge")

	return newQmlBridgeFromPointer(C.QmlBridge_NewQmlBridge(core.PointerFromQObject(parent)))
}

func (ptr *QmlBridge) DestroyQmlBridge() {
	defer qt.Recovering("QmlBridge::~QmlBridge")

	if ptr.Pointer() != nil {
		C.QmlBridge_DestroyQmlBridge(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QmlBridge) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QmlBridge::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QmlBridge) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QmlBridge::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQmlBridgeTimerEvent
func callbackQmlBridgeTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlBridge) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QmlBridge::timerEvent")

	if ptr.Pointer() != nil {
		C.QmlBridge_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QmlBridge) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QmlBridge::timerEvent")

	if ptr.Pointer() != nil {
		C.QmlBridge_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QmlBridge) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QmlBridge::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QmlBridge) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QmlBridge::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQmlBridgeChildEvent
func callbackQmlBridgeChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlBridge) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QmlBridge::childEvent")

	if ptr.Pointer() != nil {
		C.QmlBridge_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QmlBridge) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QmlBridge::childEvent")

	if ptr.Pointer() != nil {
		C.QmlBridge_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QmlBridge) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QmlBridge::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QmlBridge) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QmlBridge::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

//export callbackQmlBridgeConnectNotify
func callbackQmlBridgeConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QmlBridge::connectNotify")

	if ptr.Pointer() != nil {
		C.QmlBridge_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QmlBridge) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QmlBridge::connectNotify")

	if ptr.Pointer() != nil {
		C.QmlBridge_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QmlBridge) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QmlBridge::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QmlBridge) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QmlBridge::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQmlBridgeCustomEvent
func callbackQmlBridgeCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlBridge) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QmlBridge::customEvent")

	if ptr.Pointer() != nil {
		C.QmlBridge_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QmlBridge) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QmlBridge::customEvent")

	if ptr.Pointer() != nil {
		C.QmlBridge_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QmlBridge) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QmlBridge::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QmlBridge) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QmlBridge::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

//export callbackQmlBridgeDisconnectNotify
func callbackQmlBridgeDisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QmlBridge::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QmlBridge) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QmlBridge::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}
