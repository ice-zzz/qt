package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QIODevice struct {
	QObject
}

type QIODevice_ITF interface {
	QObject_ITF
	QIODevice_PTR() *QIODevice
}

func PointerFromQIODevice(ptr QIODevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func NewQIODeviceFromPointer(ptr unsafe.Pointer) *QIODevice {
	var n = new(QIODevice)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QIODevice_") {
		n.SetObjectName("QIODevice_" + qt.Identifier())
	}
	return n
}

func (ptr *QIODevice) QIODevice_PTR() *QIODevice {
	return ptr
}

//QIODevice::OpenModeFlag
type QIODevice__OpenModeFlag int64

const (
	QIODevice__NotOpen    = QIODevice__OpenModeFlag(0x0000)
	QIODevice__ReadOnly   = QIODevice__OpenModeFlag(0x0001)
	QIODevice__WriteOnly  = QIODevice__OpenModeFlag(0x0002)
	QIODevice__ReadWrite  = QIODevice__OpenModeFlag(QIODevice__ReadOnly | QIODevice__WriteOnly)
	QIODevice__Append     = QIODevice__OpenModeFlag(0x0004)
	QIODevice__Truncate   = QIODevice__OpenModeFlag(0x0008)
	QIODevice__Text       = QIODevice__OpenModeFlag(0x0010)
	QIODevice__Unbuffered = QIODevice__OpenModeFlag(0x0020)
)

func (ptr *QIODevice) GetChar(c string) bool {
	defer qt.Recovering("QIODevice::getChar")

	if ptr.Pointer() != nil {
		return C.QIODevice_GetChar(ptr.Pointer(), C.CString(c)) != 0
	}
	return false
}

func (ptr *QIODevice) PutChar(c string) bool {
	defer qt.Recovering("QIODevice::putChar")

	if ptr.Pointer() != nil {
		return C.QIODevice_PutChar(ptr.Pointer(), C.CString(c)) != 0
	}
	return false
}

func (ptr *QIODevice) ConnectAboutToClose(f func()) {
	defer qt.Recovering("connect QIODevice::aboutToClose")

	if ptr.Pointer() != nil {
		C.QIODevice_ConnectAboutToClose(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToClose", f)
	}
}

func (ptr *QIODevice) DisconnectAboutToClose() {
	defer qt.Recovering("disconnect QIODevice::aboutToClose")

	if ptr.Pointer() != nil {
		C.QIODevice_DisconnectAboutToClose(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToClose")
	}
}

//export callbackQIODeviceAboutToClose
func callbackQIODeviceAboutToClose(ptrName *C.char) {
	defer qt.Recovering("callback QIODevice::aboutToClose")

	var signal = qt.GetSignal(C.GoString(ptrName), "aboutToClose")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QIODevice) AtEnd() bool {
	defer qt.Recovering("QIODevice::atEnd")

	if ptr.Pointer() != nil {
		return C.QIODevice_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) BytesAvailable() int64 {
	defer qt.Recovering("QIODevice::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QIODevice_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIODevice) BytesToWrite() int64 {
	defer qt.Recovering("QIODevice::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QIODevice_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIODevice) ConnectBytesWritten(f func(bytes int64)) {
	defer qt.Recovering("connect QIODevice::bytesWritten")

	if ptr.Pointer() != nil {
		C.QIODevice_ConnectBytesWritten(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bytesWritten", f)
	}
}

func (ptr *QIODevice) DisconnectBytesWritten() {
	defer qt.Recovering("disconnect QIODevice::bytesWritten")

	if ptr.Pointer() != nil {
		C.QIODevice_DisconnectBytesWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bytesWritten")
	}
}

//export callbackQIODeviceBytesWritten
func callbackQIODeviceBytesWritten(ptrName *C.char, bytes C.longlong) {
	defer qt.Recovering("callback QIODevice::bytesWritten")

	var signal = qt.GetSignal(C.GoString(ptrName), "bytesWritten")
	if signal != nil {
		signal.(func(int64))(int64(bytes))
	}

}

func (ptr *QIODevice) CanReadLine() bool {
	defer qt.Recovering("QIODevice::canReadLine")

	if ptr.Pointer() != nil {
		return C.QIODevice_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) ConnectClose(f func()) {
	defer qt.Recovering("connect QIODevice::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QIODevice) DisconnectClose() {
	defer qt.Recovering("disconnect QIODevice::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQIODeviceClose
func callbackQIODeviceClose(ptrName *C.char) bool {
	defer qt.Recovering("callback QIODevice::close")

	var signal = qt.GetSignal(C.GoString(ptrName), "close")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QIODevice) ErrorString() string {
	defer qt.Recovering("QIODevice::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QIODevice_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIODevice) IsOpen() bool {
	defer qt.Recovering("QIODevice::isOpen")

	if ptr.Pointer() != nil {
		return C.QIODevice_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) IsReadable() bool {
	defer qt.Recovering("QIODevice::isReadable")

	if ptr.Pointer() != nil {
		return C.QIODevice_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) IsSequential() bool {
	defer qt.Recovering("QIODevice::isSequential")

	if ptr.Pointer() != nil {
		return C.QIODevice_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) IsTextModeEnabled() bool {
	defer qt.Recovering("QIODevice::isTextModeEnabled")

	if ptr.Pointer() != nil {
		return C.QIODevice_IsTextModeEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) IsWritable() bool {
	defer qt.Recovering("QIODevice::isWritable")

	if ptr.Pointer() != nil {
		return C.QIODevice_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) Open(mode QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QIODevice::open")

	if ptr.Pointer() != nil {
		return C.QIODevice_Open(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QIODevice) OpenMode() QIODevice__OpenModeFlag {
	defer qt.Recovering("QIODevice::openMode")

	if ptr.Pointer() != nil {
		return QIODevice__OpenModeFlag(C.QIODevice_OpenMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIODevice) Peek2(maxSize int64) *QByteArray {
	defer qt.Recovering("QIODevice::peek")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QIODevice_Peek2(ptr.Pointer(), C.longlong(maxSize)))
	}
	return nil
}

func (ptr *QIODevice) Peek(data string, maxSize int64) int64 {
	defer qt.Recovering("QIODevice::peek")

	if ptr.Pointer() != nil {
		return int64(C.QIODevice_Peek(ptr.Pointer(), C.CString(data), C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QIODevice) Pos() int64 {
	defer qt.Recovering("QIODevice::pos")

	if ptr.Pointer() != nil {
		return int64(C.QIODevice_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIODevice) Read2(maxSize int64) *QByteArray {
	defer qt.Recovering("QIODevice::read")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QIODevice_Read2(ptr.Pointer(), C.longlong(maxSize)))
	}
	return nil
}

func (ptr *QIODevice) Read(data string, maxSize int64) int64 {
	defer qt.Recovering("QIODevice::read")

	if ptr.Pointer() != nil {
		return int64(C.QIODevice_Read(ptr.Pointer(), C.CString(data), C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QIODevice) ReadAll() *QByteArray {
	defer qt.Recovering("QIODevice::readAll")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QIODevice_ReadAll(ptr.Pointer()))
	}
	return nil
}

func (ptr *QIODevice) ConnectReadChannelFinished(f func()) {
	defer qt.Recovering("connect QIODevice::readChannelFinished")

	if ptr.Pointer() != nil {
		C.QIODevice_ConnectReadChannelFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readChannelFinished", f)
	}
}

func (ptr *QIODevice) DisconnectReadChannelFinished() {
	defer qt.Recovering("disconnect QIODevice::readChannelFinished")

	if ptr.Pointer() != nil {
		C.QIODevice_DisconnectReadChannelFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readChannelFinished")
	}
}

//export callbackQIODeviceReadChannelFinished
func callbackQIODeviceReadChannelFinished(ptrName *C.char) {
	defer qt.Recovering("callback QIODevice::readChannelFinished")

	var signal = qt.GetSignal(C.GoString(ptrName), "readChannelFinished")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QIODevice) ReadLine2(maxSize int64) *QByteArray {
	defer qt.Recovering("QIODevice::readLine")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QIODevice_ReadLine2(ptr.Pointer(), C.longlong(maxSize)))
	}
	return nil
}

func (ptr *QIODevice) ReadLine(data string, maxSize int64) int64 {
	defer qt.Recovering("QIODevice::readLine")

	if ptr.Pointer() != nil {
		return int64(C.QIODevice_ReadLine(ptr.Pointer(), C.CString(data), C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QIODevice) ConnectReadyRead(f func()) {
	defer qt.Recovering("connect QIODevice::readyRead")

	if ptr.Pointer() != nil {
		C.QIODevice_ConnectReadyRead(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyRead", f)
	}
}

func (ptr *QIODevice) DisconnectReadyRead() {
	defer qt.Recovering("disconnect QIODevice::readyRead")

	if ptr.Pointer() != nil {
		C.QIODevice_DisconnectReadyRead(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyRead")
	}
}

//export callbackQIODeviceReadyRead
func callbackQIODeviceReadyRead(ptrName *C.char) {
	defer qt.Recovering("callback QIODevice::readyRead")

	var signal = qt.GetSignal(C.GoString(ptrName), "readyRead")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QIODevice) Reset() bool {
	defer qt.Recovering("QIODevice::reset")

	if ptr.Pointer() != nil {
		return C.QIODevice_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) Seek(pos int64) bool {
	defer qt.Recovering("QIODevice::seek")

	if ptr.Pointer() != nil {
		return C.QIODevice_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QIODevice) SetTextModeEnabled(enabled bool) {
	defer qt.Recovering("QIODevice::setTextModeEnabled")

	if ptr.Pointer() != nil {
		C.QIODevice_SetTextModeEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QIODevice) Size() int64 {
	defer qt.Recovering("QIODevice::size")

	if ptr.Pointer() != nil {
		return int64(C.QIODevice_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIODevice) UngetChar(c string) {
	defer qt.Recovering("QIODevice::ungetChar")

	if ptr.Pointer() != nil {
		C.QIODevice_UngetChar(ptr.Pointer(), C.CString(c))
	}
}

func (ptr *QIODevice) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QIODevice::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QIODevice_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QIODevice) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QIODevice::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QIODevice_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QIODevice) Write3(byteArray QByteArray_ITF) int64 {
	defer qt.Recovering("QIODevice::write")

	if ptr.Pointer() != nil {
		return int64(C.QIODevice_Write3(ptr.Pointer(), PointerFromQByteArray(byteArray)))
	}
	return 0
}

func (ptr *QIODevice) Write2(data string) int64 {
	defer qt.Recovering("QIODevice::write")

	if ptr.Pointer() != nil {
		return int64(C.QIODevice_Write2(ptr.Pointer(), C.CString(data)))
	}
	return 0
}

func (ptr *QIODevice) Write(data string, maxSize int64) int64 {
	defer qt.Recovering("QIODevice::write")

	if ptr.Pointer() != nil {
		return int64(C.QIODevice_Write(ptr.Pointer(), C.CString(data), C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QIODevice) DestroyQIODevice() {
	defer qt.Recovering("QIODevice::~QIODevice")

	if ptr.Pointer() != nil {
		C.QIODevice_DestroyQIODevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}