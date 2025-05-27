package cmd

import (
	"fmt"
	"log"
	"os"
	"strings" // Додаємо для роботи з рядками
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	Teletoken = os.Getenv("TELE_TOKEN")
)

// tgbotCmd represents the tgbot command
var tgbotCmd = &cobra.Command{
	Use:     "tgbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tgbot %s started\n", appVersion) // Додав \n для кращого відображення
		tgbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  Teletoken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		}

		// --- Обробник для команди /start ---
		tgbot.Handle("/start", func(m telebot.Context) error {
			log.Printf("Received /start from %s", m.Sender().Username)
			return m.Send(fmt.Sprintf("Привіт, %s! Я ваш бот. Для допомоги використовуйте команду /help.", m.Sender().FirstName))
		})

		// --- Обробник для команди /help ---
		tgbot.Handle("/help", func(m telebot.Context) error {
			log.Printf("Received /help from %s", m.Sender().Username)
			helpMessage := "Доступні команди:\n" +
				"/start - Почати роботу з ботом\n" +
				"/help - Показати цю довідку\n" +
				"/echo <текст> - Я повторю ваш текст\n" +
				"/wordcount <текст> - Порахувати кількість слів у вашому тексті\n" +
				"Будь-який інший текст - Я відповім, що не розумію команду."
			return m.Send(helpMessage)
		})

		// --- Обробник для команди /echo ---
		tgbot.Handle("/echo", func(m telebot.Context) error {
			text := m.Text() // Отримуємо весь текст повідомлення
			// Якщо текст починається з "/echo ", то беремо все, що після пробілу
			if strings.HasPrefix(text, "/echo ") {
				echoText := strings.TrimPrefix(text, "/echo ")
				log.Printf("Received /echo command with text: %s", echoText)
				return m.Send(fmt.Sprintf("Ви сказали: %s", echoText))
			}
			log.Printf("Received /echo command without text from %s", m.Sender().Username)
			return m.Send("Будь ласка, введіть текст після команди /echo, наприклад: /echo Привіт!")
		})

		// --- Обробник для команди /wordcount ---
		tgbot.Handle("/wordcount", func(m telebot.Context) error {
			text := m.Text()
			if strings.HasPrefix(text, "/wordcount ") {
				input := strings.TrimPrefix(text, "/wordcount ")
				words := strings.Fields(input) // Розбиває рядок на слова за пробілами
				count := len(words)
				log.Printf("Received /wordcount command with text: '%s'. Words count: %d", input, count)
				return m.Send(fmt.Sprintf("У вашому тексті %d слів.", count))
			}
			log.Printf("Received /wordcount command without text from %s", m.Sender().Username)
			return m.Send("Будь ласка, введіть текст після команди /wordcount, наприклад: /wordcount Привіт, світ!")
		})

		// --- Обробник будь-якого текстового повідомлення, якщо воно не є командою ---
		tgbot.Handle(telebot.OnText, func(m telebot.Context) error {
			// Перевіряємо, чи це не одна зі стандартних команд
			if strings.HasPrefix(m.Text(), "/") {
				// Якщо це команда, але не оброблена вище, то повідомляємо про невідому команду
				log.Printf("Received unknown command: %s from %s", m.Text(), m.Sender().Username)
				return m.Send("Невідома команда. Спробуйте /help для списку команд.")
			}

			// Якщо це просто текст, і не команда, то можемо відповісти "Я вас не розумію"
			log.Printf("Received unhandled text: %s from %s", m.Text(), m.Sender().Username)
			return m.Send("Я отримав ваше повідомлення, але не розумію, що мені з ним робити. Спробуйте використати команди.")
		})

		log.Println("Bot is starting...")
		tgbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(tgbotCmd)
}
