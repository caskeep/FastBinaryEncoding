// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.2.0.0

package proto

import "errors"
import "../fbe"

// Workaround for Go unused imports issue
var _ = fbe.Version

// Fast Binary Encoding proto final receiver
type FinalReceiver struct {
    *fbe.Receiver
    orderValue *Order
    orderModel *OrderFinalModel
    balanceValue *Balance
    balanceModel *BalanceFinalModel
    accountValue *Account
    accountModel *AccountFinalModel

    // Receive Order handler
    HandlerOnReceiveOrder OnReceiveOrder
    // Receive Balance handler
    HandlerOnReceiveBalance OnReceiveBalance
    // Receive Account handler
    HandlerOnReceiveAccount OnReceiveAccount
}

// Create a new proto final receiver with an empty buffer
func NewFinalReceiver() *FinalReceiver {
    return NewFinalReceiverWithBuffer(fbe.NewEmptyBuffer())
}

// Create a new proto final receiver with the given buffer
func NewFinalReceiverWithBuffer(buffer *fbe.Buffer) *FinalReceiver {
    receiver := &FinalReceiver{
        fbe.NewReceiver(buffer, true),
        NewOrder(),
        NewOrderFinalModel(buffer),
        NewBalance(),
        NewBalanceFinalModel(buffer),
        NewAccount(),
        NewAccountFinalModel(buffer),
        nil,
        nil,
        nil,
    }
    receiver.SetupHandlerOnReceive(receiver)
    receiver.SetupHandlerOnReceiveOrderFunc(func(value *Order) {})
    receiver.SetupHandlerOnReceiveBalanceFunc(func(value *Balance) {})
    receiver.SetupHandlerOnReceiveAccountFunc(func(value *Account) {})
    return receiver
}

// Setup handlers
func (r *FinalReceiver) SetupHandlers(handlers interface{}) {
    r.Receiver.SetupHandlers(handlers)
    if handler, ok := handlers.(OnReceiveOrder); ok {
        r.SetupHandlerOnReceiveOrder(handler)
    }
    if handler, ok := handlers.(OnReceiveBalance); ok {
        r.SetupHandlerOnReceiveBalance(handler)
    }
    if handler, ok := handlers.(OnReceiveAccount); ok {
        r.SetupHandlerOnReceiveAccount(handler)
    }
}

// Setup receive Order handler
func (r *FinalReceiver) SetupHandlerOnReceiveOrder(handler OnReceiveOrder) { r.HandlerOnReceiveOrder = handler }
// Setup receive Order handler function
func (r *FinalReceiver) SetupHandlerOnReceiveOrderFunc(function func(value *Order)) { r.HandlerOnReceiveOrder = OnReceiveOrderFunc(function) }
// Setup receive Balance handler
func (r *FinalReceiver) SetupHandlerOnReceiveBalance(handler OnReceiveBalance) { r.HandlerOnReceiveBalance = handler }
// Setup receive Balance handler function
func (r *FinalReceiver) SetupHandlerOnReceiveBalanceFunc(function func(value *Balance)) { r.HandlerOnReceiveBalance = OnReceiveBalanceFunc(function) }
// Setup receive Account handler
func (r *FinalReceiver) SetupHandlerOnReceiveAccount(handler OnReceiveAccount) { r.HandlerOnReceiveAccount = handler }
// Setup receive Account handler function
func (r *FinalReceiver) SetupHandlerOnReceiveAccountFunc(function func(value *Account)) { r.HandlerOnReceiveAccount = OnReceiveAccountFunc(function) }

// Receive message handler
func (r *FinalReceiver) OnReceive(fbeType int, buffer []byte) (bool, error) {
    switch fbeType {
    case r.orderModel.FBEType():
        // Deserialize the value from the FBE stream
        r.orderModel.Buffer().Attach(buffer)
        if !r.orderModel.Verify() {
            return false, errors.New("proto.Order validation failed")
        }
        deserialized, err := r.orderModel.DeserializeValue(r.orderValue)
        if deserialized <= 0 {
            return false, errors.New("proto.Order deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.orderValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveOrder.OnReceiveOrder(r.orderValue)
        return true, nil
    case r.balanceModel.FBEType():
        // Deserialize the value from the FBE stream
        r.balanceModel.Buffer().Attach(buffer)
        if !r.balanceModel.Verify() {
            return false, errors.New("proto.Balance validation failed")
        }
        deserialized, err := r.balanceModel.DeserializeValue(r.balanceValue)
        if deserialized <= 0 {
            return false, errors.New("proto.Balance deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.balanceValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveBalance.OnReceiveBalance(r.balanceValue)
        return true, nil
    case r.accountModel.FBEType():
        // Deserialize the value from the FBE stream
        r.accountModel.Buffer().Attach(buffer)
        if !r.accountModel.Verify() {
            return false, errors.New("proto.Account validation failed")
        }
        deserialized, err := r.accountModel.DeserializeValue(r.accountValue)
        if deserialized <= 0 {
            return false, errors.New("proto.Account deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.accountValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveAccount.OnReceiveAccount(r.accountValue)
        return true, nil
    }

    return false, nil
}
