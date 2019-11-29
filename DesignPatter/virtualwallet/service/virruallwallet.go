package service

type VirtuallWallet struct {
	id         int64
	createTime int64
	balnace    float64
}

func (vw *VirtuallWallet) Balnace() float64 {
	return vw.balnace
}

func (vw *VirtuallWallet) Debit(amount float64) {
	return vw.balnace.subtract(amount)

}

func (vw *VirtuallWallet) Credit(amount float64) {
	return vw.balnace.add(amount)
}
