package ai

import (
	"context"
	"log"
	"time"
)

// MockAIProvider имитирует работу AI модели для тестирования
type MockAIProvider struct {
	counter int
}

func NewMockAIProvider() *MockAIProvider {
	return &MockAIProvider{counter: 0}
}

func (m *MockAIProvider) Analyze(ctx context.Context, input AnalysisInput) (AnalysisOutput, error) {
	m.counter++

	// Генерируем разные подсказки в зависимости от типа и содержимого
	var hint string
	var tasks []string
	var warnings []string

	if input.Type == "audio" {
		// Анализируем аудиотранскрипцию
		if contains(input.TranscriptText, "CPU") {
			hint = "⚠️ Критическая нагрузка CPU обнаружена! Рекомендуй проверить top -H и recent deploys."
			tasks = []string{
				"Проверить использование процессов: top -H",
				"Просмотреть последние деплои",
				"Проверить memory leaks",
			}
			warnings = []string{"Возможен DDoS или infinite loop"}
		} else if contains(input.TranscriptText, "Memory leak") {
			hint = "🔴 Memory leak обнаружен! Выполни откат к предыдущему коммиту и проверь heap dump."
			tasks = []string{
				"Выполнить git rollback",
				"Проверить heap dump",
				"Запустить memory profiler",
			}
			warnings = []string{"Откат может потребовать перезагрузку", "Проверь dependencies перед откатом"}
		} else if contains(input.TranscriptText, "откатиться") {
			hint = "🔄 Перед откатом убедитесь, что завершены все ongoing transactions."
			tasks = []string{
				"Проверить active transactions",
				"Создать backup текущего состояния",
				"Выполнить rollback",
			}
			warnings = []string{"Может потребоваться downtime", "Уведомить stakeholders"}
		} else if contains(input.TranscriptText, "Database") {
			hint = "💾 Database issue! Проверь connection pool и query performance."
			tasks = []string{
				"Проверить connection pool: SHOW PROCESSLIST",
				"Анализировать slow queries",
				"Проверить disk space",
			}
			warnings = []string{"Возможна блокировка", "Проверь репликацию"}
		} else if contains(input.TranscriptText, "API") {
			hint = "🌐 API issue! Проверь endpoints availability и load balancer."
			tasks = []string{
				"Проверить API health endpoints",
				"Посмотреть логи load balancer'а",
				"Проверить rate limits",
			}
			warnings = []string{"Возможен rate limiting", "Проверь DNS resolution"}
		} else {
			hint = "ℹ️ Подсказка: используй стандартные инструменты диагностики (top, netstat, curl)."
			tasks = []string{
				"Собрать метрики системы",
				"Проверить логи приложения",
				"Провести базовую диагностику",
			}
			warnings = []string{"Собери достаточно контекста перед действиями"}
		}
	} else if input.Type == "vision" {
		// Анализируем текст из скриншотов
		if contains(input.OCRText, "ERROR") || contains(input.OCRText, "error") {
			hint = "🔴 Обнаружена ERROR в логах! Определи компонент и проверь stack trace."
			tasks = []string{
				"Найти точный компонент ошибки",
				"Посмотреть полный stack trace",
				"Проверить related errors",
			}
			warnings = []string{"Может быть cascade failure"}
		} else if contains(input.OCRText, "95%") || contains(input.OCRText, "CPU") {
			hint = "🔥 Критическая метрика обнаружена: CPU/Memory на уровне 95%+. Срочно проверь top!"
			tasks = []string{
				"Выполнить: top -H | head -20",
				"Посмотреть процессы с максимальным usage",
				"Проверить network I/O",
			}
			warnings = []string{"Система может скоро упасть", "Приготовься к аварийному откату"}
		} else if contains(input.OCRText, "CrashLoopBackOff") {
			hint = "💥 Pod находится в CrashLoopBackOff! Проверь последние логи и events."
			tasks = []string{
				"Просмотреть pod events: kubectl describe pod",
				"Посмотреть логи: kubectl logs",
				"Проверить resource limits",
			}
			warnings = []string{"Возможна нехватка ресурсов", "Проверь health checks"}
		} else if contains(input.OCRText, "503") || contains(input.OCRText, "Unavailable") {
			hint = "⚠️ Service Unavailable (503). Проверь backend status и load balancer configuration."
			tasks = []string{
				"Проверить health endpoints всех backend'ов",
				"Посмотреть логи load balancer'а",
				"Проверить network connectivity",
			}
			warnings = []string{"Возможен cascading failure"}
		} else {
			hint = "📊 Метрики собраны. Проанализируй тренды и сравни с baseline'ом."
			tasks = []string{
				"Определить anomalies",
				"Сравнить с историческими данными",
				"Проверить correlation с recent changes",
			}
			warnings = []string{}
		}
	}

	log.Printf("🤖 Mock AI #%d (type=%s): %s", m.counter, input.Type, hint)

	// Симулируем задержку AI анализа
	select {
	case <-time.After(1000 * time.Millisecond):
		return AnalysisOutput{
			Hint:       hint,
			Tasks:      tasks,
			Warnings:   warnings,
			Confidence: 0.85,
		}, nil
	case <-ctx.Done():
		return AnalysisOutput{}, ctx.Err()
	}
}

func (m *MockAIProvider) Health(ctx context.Context) error {
	log.Println("🤖 Mock AI Provider is healthy")
	return nil
}

func contains(s, substr string) bool {
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
