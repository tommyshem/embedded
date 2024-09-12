// bare bones blink led on esp32 dev board
#define GPIO_OUT_REG 0x3FF44004
#define GPIO_OUT_ENABLE 0x3FF44020
#define LED_PIN 2
typedef unsigned int            u_int32_t;

void delay_ms(u_int32_t ms){
    for (u_int32_t i = 0; i < ms; i++)
    {
        for (u_int32_t i = 0; i < 3000; i++)
        {
            volatile u_int32_t temp = 0;
        }
        
    }
    
}

void app_main(void)
{
volatile u_int32_t* gpio_out =(volatile u_int32_t*)GPIO_OUT_REG;
volatile u_int32_t* gpio_enable = (volatile u_int32_t*)GPIO_OUT_ENABLE;
*gpio_enable = (1<<LED_PIN);  // setup pin 2 as an output

while(1){
    *gpio_out ^=(1<<LED_PIN);  // toggle on/ off with XOR
    delay_ms(1000); // delay 1 second
}
}
