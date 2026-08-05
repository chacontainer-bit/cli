import { useEffect, useRef, useState } from 'react'
import { useLocation } from 'react-router-dom'
import { Bot, Send, X, Sparkles } from 'lucide-react'
import { answerQuery, generateInsights, getGreeting } from '../lib/assistant'

function openingMessage() {
  const insights = generateInsights()
  const critical = insights.filter(i => i.severity === 'critico')
  if (critical.length > 0) {
    return `${getGreeting()}, Humberto. Antes de que preguntes: ${critical[0].title.toLowerCase()}. ${critical[0].message}`
  }
  if (insights.length > 0) {
    return `${getGreeting()}, Humberto. Nada crítico, pero tengo ${insights.length} cosa(s) que vale la pena que veas cuando puedas.`
  }
  return `${getGreeting()}, Humberto. Todo dentro de parámetros normales. Pregúntame por flota, proyectos, inventario o talleres.`
}

export default function AssistantWidget() {
  const location = useLocation()
  const [open, setOpen] = useState(false)
  const [messages, setMessages] = useState(() => [{ role: 'assistant', text: openingMessage() }])
  const [input, setInput] = useState('')
  const scrollRef = useRef(null)

  useEffect(() => {
    if (scrollRef.current) scrollRef.current.scrollTop = scrollRef.current.scrollHeight
  }, [messages, open])

  if (location.pathname === '/asistente') return null

  function handleSend(e) {
    e.preventDefault()
    const text = input.trim()
    if (!text) return
    const reply = answerQuery(text)
    setMessages(prev => [...prev, { role: 'user', text }, { role: 'assistant', text: reply }])
    setInput('')
  }

  return (
    <div className="fixed bottom-5 right-5 left-5 sm:left-auto z-50 flex flex-col items-end">
      {open && (
        <div className="mb-3 w-full sm:w-96 max-w-full h-[28rem] bg-[#111827] border border-[#1f2937] rounded-xl shadow-2xl flex flex-col overflow-hidden">
          <div className="px-4 py-3 border-b border-[#1f2937] flex items-center justify-between bg-[#0a0f1a]">
            <div className="flex items-center gap-2">
              <div className="w-7 h-7 rounded-full bg-[#f97316]/10 flex items-center justify-center">
                <Bot size={15} className="text-[#f97316]" />
              </div>
              <div>
                <p className="text-[#f9fafb] text-xs font-semibold">Asistente CHACONTAINER</p>
                <p className="text-[#6b7280] text-[10px]">Panel ejecutivo</p>
              </div>
            </div>
            <button onClick={() => setOpen(false)} className="text-[#6b7280] hover:text-[#f9fafb]">
              <X size={16} />
            </button>
          </div>
          <div ref={scrollRef} className="flex-1 overflow-y-auto p-3 space-y-2">
            {messages.map((m, i) => (
              <div key={i} className={`flex ${m.role === 'user' ? 'justify-end' : 'justify-start'}`}>
                <div
                  className={`max-w-[85%] rounded-lg px-3 py-2 text-xs leading-relaxed whitespace-pre-line ${
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
          <form onSubmit={handleSend} className="p-2 border-t border-[#1f2937] flex items-center gap-2">
            <input
              value={input}
              onChange={e => setInput(e.target.value)}
              placeholder="Pregunta por flota, proyectos, inventario..."
              className="flex-1 bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-xs text-[#f9fafb] placeholder-[#6b7280] focus:outline-none focus:border-[#f97316]"
            />
            <button
              type="submit"
              className="w-8 h-8 flex-shrink-0 rounded-lg bg-[#f97316] hover:bg-[#ea6a0e] flex items-center justify-center transition-colors"
            >
              <Send size={14} className="text-white" />
            </button>
          </form>
        </div>
      )}
      <button
        onClick={() => setOpen(v => !v)}
        className="w-14 h-14 rounded-full bg-[#f97316] hover:bg-[#ea6a0e] shadow-xl flex items-center justify-center transition-colors relative"
      >
        {open ? <X size={22} className="text-white" /> : <Sparkles size={22} className="text-white" />}
      </button>
    </div>
  )
}
