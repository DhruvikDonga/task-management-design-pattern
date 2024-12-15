package notify

type Notify interface {
	Update(message string)
}

type NotifcationReceivers interface {
	RegisterReceiver(Notify)
	//UnRegisterReceiver(Notify)
	NotifyRegisters(message string)
}
