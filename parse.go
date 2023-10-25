package main

const (
	// ビット2: Primary Account Number
	BitPrimaryAccountNumber = 0x4000000000000000
	// ビット3: Processing Code
	BitProcessingCode = 0x2000000000000000
	// ビット4: Amount Transaction
	BitAmountTransaction = 0x1000000000000000
	// ビット14: Expiration Date
	BitExpirationDate = 0x0004000000000000
)

// ISO 8583 メッセージ構造体
type Message struct {
	// メッセージ: 3桁の数字
	// ISO 8583 バージョンの 0 は含まないことに注意してください
	Type uint16

	// カード会員番号 (PAN): 16桁の数字
	PrimaryAccountNumber uint64

	// 処理コード: 6桁の数字
	ProcessingCode uint32

	// 取引金額: 12桁の数字
	AmountTransaction uint64

	// 有効期限: YYMM の4桁の数字
	ExpirationDate uint16
}

// ISO 8583 メッセージのバイト列をパースして Message を返す関数
func Parse(buf []byte) *Message {
	var msg Message

	// TODO: バイト列 buf をパースして msg の各フィールドに値を設定してください

	return &msg
}
