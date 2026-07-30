import { useEffect, useRef, useState } from 'react'
import { Bot, Send, AlertTriangle } from 'lucide-react'
import { answerQuery, generateInsights, getGreeting } from '../lib/assistant'

const severityConfig = {
  critico: { border: 'border-red-500', bg: 'bg-red-500/10', dot: 'bg-red-500', label: 'CRÍTICO', text: 'text-red-400' },
  advertencia: { border: 'border-amber-500', bg: 'bg-amber-500/10', dot: 'bg-amber-500', label: 'AVISO', text: 'text-amber-400' },
  info: { border: 'border-blue-500', bg: 'bg-blue-500/10', dot: 'bg-blue-500', label: 'INFO', text: 'text-blue-400' },
}

function openingMessage(insights) {
  const critical = insights.filter(i => i.severity === 'critico')
  if (critical.length > 0) {
    return `${getGreeting()}, Humberto. Vamos directo al punto: ${critical.length} tema(s) crítico(s) requieren tu atención hoy. Los tienes en el panel de la derecha, o pregúntame por cualquier área — flota, proyectos, inventario, talleres.`
  }
  if (insights.length > 0) {
    return `${getGreeting()}, Humberto. Nada crítico ahora mismo, pero hay ${insights.length} cosa(s) que conviene que veas. Están en el panel de la derecha, ordenadas por prioridad.`
  }
  return `${getGreeting()}, Humberto. Todo dentro de parámetros normales en flota, proyectos, inventario y talleres. Pregúntame lo que necesites.`
}

export default function Asistente() {
  const [insights] = useState(() => generateInsights())
  const [messages, setMessages] = useState(() => [{ role: 'assistant', text: openingMessage(insights) }])
  const [input, setInput] = useState('')
  const scrollRef = useRef(null)

  useEffect(() => {
    if (scrollRef.current) scrollRef.current.scrollTop = scrollRef.current.scrollHeight
  }, [messages])

  function handleSend(e) {
    e.preventDefault()
    const text = input.trim()
    if (!text) return
    const reply = answerQuery(text)
    setMessages(prev => [...prev, { role: 'user', text }, { role: 'assistant', text: reply }])
    setInput('')
  }

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-xl font-bold text-[#f9fafb] flex items-center gap-2">
          <Bot className="text-[#f97316]" size={22} />
          Asistente CHACONTAINER
        </h1>
        <p className="text-[#6b7280] text-sm mt-0.5">Panel de control ejecutivo — visión completa de la operación, para ti.</p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
        {/* Chat */}
        <div className="lg:col-span-2 bg-[#111827] border border-[#1f2937] rounded-xl flex flex-col h-[32rem]">
          <div ref={scrollRef} className="flex-1 overflow-y-auto p-4 space-y-3">
            {messages.map((m, i) => (
              <div key={i} className={`flex ${m.role === 'user' ? 'justify-end' : 'justify-start'}`}>
                <div
                  className={`max-w-[80%] rounded-lg px-4 py-2.5 text-sm leading-relaxed whitespace-pre-line ${
                    m.role === 'user'
                      ? 'bg-[#f97316] text-white'
                      : 'bg-[#1a2235] text-[#d1d5db] border border-[#1f2937]'
                  }`}
                >
                  {m.text}
                </div>
              </div>
            ))}
          </div>
          <form onSubmit={handleSend} className="p-3 border-t border-[#1f2937] flex items-center gap-2">
            <input
              value={input}
              onChange={e => setInput(e.target.value)}
              placeholder="Pregunta por flota, alertas, proyectos, inventario o talleres..."
              className="flex-1 bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2.5 text-sm text-[#f9fafb] placeholder-[#6b7280] focus:outline-none focus:border-[#f97316]"
            />
            <button
              type="submit"
              className="px-4 h-10 flex-shrink-0 rounded-lg bg-[#f97316] hover:bg-[#ea6a0e] flex items-center gap-2 text-sm text-white font-medium transition-colors"
            >
              <Send size={15} />
              Enviar
            </button>
          </form>
        </div>

        {/* Insights feed */}
        <div className="bg-[#111827] border border-[#1f2937] rounded-xl flex flex-col h-[32rem]">
          <div className="px-4 py-3 border-b border-[#1f2937] flex items-center justify-between">
            <h2 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2">
              <AlertTriangle size={15} className="text-[#f97316]" />
              Lo que anticipé
            </h2>
            <span className="text-xs font-mono text-[#6b7280]">{insights.length} total</span>
          </div>
          <div className="flex-1 overflow-y-auto divide-y divide-[#1f2937]">
            {insights.length === 0 && (
              <p className="text-[#6b7280] text-sm p-4">Sin nada que reportar por ahora.</p>
            )}
            {insights.map(item => {
              const cfg = severityConfig[item.severity]
              return (
                <div key={item.id} className={`flex gap-3 p-3 border-l-2 ${cfg.border} ${cfg.bg}`}>
                  <div className={`w-1.5 h-1.5 rounded-full mt-1.5 flex-shrink-0 ${cfg.dot}`} />
                  <div className="min-w-0 flex-1">
                    <div className="flex items-center gap-2 mb-0.5">
                      <span className={`text-xs font-bold font-mono ${cfg.text}`}>{cfg.label}</span>
                      <span className="text-[#6b7280] text-xs">{item.module}</span>
                    </div>
                    <p className="text-[#f9fafb] text-xs font-medium">{item.title}</p>
                    <p className="text-[#d1d5db] text-xs leading-relaxed mt-0.5">{item.message}</p>
                  </div>
                </div>
              )
            })}
          </div>
        </div>
      </div>
    </div>
  )
}
