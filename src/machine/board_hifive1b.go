// +build hifive1b

package machine

const (
	D0  = P16
	D1  = P17
	D2  = P18
	D3  = P19 // Green LED/PWM (PWM1_PWM1)
	D4  = P20 // PWM (PWM1_PWM0)
	D5  = P21 // Blue LED/PWM (PWM1_PWM2)
	D6  = P22 // Red LED/PWM (PWM1_PWM3)
	D7  = P16
	D8  = NoPin // PWM?
	D9  = P01
	D10 = P02   // SPI1_CS0
	D11 = P03   // SPI1_DQ0 (MOSI)
	D12 = P04   // SPI1_DQ1 (MISO)
	D13 = P05   // SPI1_SCK
	D14 = NoPin // not connected
	D15 = P09   // does not seem to work?
	D16 = P10   // PWM (PWM2_PWM0)
	D17 = P11   // PWM (PWM2_PWM1)
	D18 = P12   // SDA (I2C0_SDA)/PWM (PWM2_PWM2)
	D19 = P13   // SDL (I2C0_SCL)/PWM (PWM2_PWM3)
)

const (
	LED       = LED1
	LED1      = LED_RED
	LED2      = LED_GREEN
	LED3      = LED_BLUE
	LED_RED   = P22
	LED_GREEN = P19
	LED_BLUE  = P21
)

const (
	// TODO: figure out the pin numbers for these.
	UART_TX_PIN = NoPin
	UART_RX_PIN = NoPin
)

// SPI pins
const (
	SPI0_SCK_PIN  = NoPin
	SPI0_MOSI_PIN = NoPin
	SPI0_MISO_PIN = NoPin

	SPI1_SCK_PIN  = D13
	SPI1_MOSI_PIN = D11
	SPI1_MISO_PIN = D12
)
