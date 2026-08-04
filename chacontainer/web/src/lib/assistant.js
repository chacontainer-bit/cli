import { containers, products, projects, workshops, kpis } from '../data/mockData'

const LOW_STOCK_THRESHOLD = 10
const WASH_RETIREMENT_THRESHOLD = 25
const BUDGET_WARNING_RATIO = 0.85
const WORKSHOP_BUSY_RATIO = 0.75

const WAREHOUSE_LABELS = {
  bodega_edomex: 'Bodega Estado de México',
  bodega_queretaro: 'Bodega Querétaro',
  bodega_puebla: 'Bodega Puebla',
}

export function getGreeting(date = new Date()) {
  const hour = date.getHours()
  if (hour < 12) return 'Buenos días'
  if (hour < 19) return 'Buenas tardes'
  return 'Buenas noches'
}

let insightSeq = 0
function insight(severity, title, message, module) {
  insightSeq += 1
  return { id: `INS-${insightSeq}`, severity, title, message, module }
}

// Cross-module analysis over the mock data — this is the "anticipates before you ask" layer.
export function generateInsights() {
  const items = []

  const perdidos = containers.filter(c => c.status === 'perdido')
  perdidos.forEach(c => {
    items.push(insight(
      'critico',
      `${c.id} sin ubicación confirmada`,
      `Último evento registrado: ${c.lastEvent}. Cliente asociado: ${c.client || 'sin cliente'}. Recomiendo abrir caso de pérdida si no aparece en las próximas 24h.`,
      'trazabilidad',
    ))
  })

  const noConformes = containers.filter(c => c.status === 'no_conforme')
  if (noConformes.length > 0) {
    items.push(insight(
      'advertencia',
      `${noConformes.length} contenedor(es) no conforme(s)`,
      `${noConformes.map(c => c.id).join(', ')}. Ninguno debería volver a circular sin pasar por reparación y control de calidad.`,
      'trazabilidad',
    ))
  }

  const paraRetiro = containers.filter(c => c.washes > WASH_RETIREMENT_THRESHOLD)
  if (paraRetiro.length > 0) {
    items.push(insight(
      'info',
      `${paraRetiro.length} contenedor(es) cerca de fin de vida útil`,
      `${paraRetiro.map(c => `${c.id} (${c.washes} lavados)`).join(', ')}. Vale la pena evaluarlos para retiro o reacondicionamiento antes de que fallen en operación.`,
      'trazabilidad',
    ))
  }

  Object.entries(WAREHOUSE_LABELS).forEach(([key, label]) => {
    products.forEach(p => {
      const stock = p.stock[key]
      if (stock !== undefined && stock < LOW_STOCK_THRESHOLD) {
        items.push(insight(
          'advertencia',
          `Stock bajo: ${p.name}`,
          `${label} tiene ${stock} unidades, por debajo del mínimo recomendado. Pedido mínimo del proveedor: ${p.minOrder}.`,
          'marketplace',
        ))
      }
    })
  })

  const riesgoAlto = projects.filter(p => p.risk === 'alto' && p.status !== 'cerrado')
  riesgoAlto.forEach(p => {
    items.push(insight(
      'advertencia',
      `${p.id} en riesgo alto`,
      `${p.name} (${p.client}) — fase ${p.phase}/${p.totalPhases}. Necesita revisión antes de seguir avanzando.`,
      'proyectos',
    ))
  })

  const presupuestoAjustado = projects.filter(p =>
    p.status === 'en_curso' && p.budget > 0 && (p.spent / p.budget) >= BUDGET_WARNING_RATIO && p.phase < p.totalPhases
  )
  presupuestoAjustado.forEach(p => {
    const pct = Math.round((p.spent / p.budget) * 100)
    items.push(insight(
      'advertencia',
      `${p.id} con presupuesto ajustado`,
      `${p.name} lleva ${pct}% del presupuesto ejecutado y aún le faltan ${p.totalPhases - p.phase} fase(s). Conviene revisar el forecast con ${p.manager}.`,
      'proyectos',
    ))
  })

  const propuestas = projects.filter(p => p.status === 'propuesta')
  if (propuestas.length > 0) {
    items.push(insight(
      'info',
      `${propuestas.length} proyecto(s) esperando decisión`,
      `${propuestas.map(p => p.id).join(', ')} siguen como propuesta. Cada día sin decidir es tiempo que se le resta a la fase de implementación.`,
      'proyectos',
    ))
  }

  const talleresOcupados = workshops.filter(w => w.current / w.capacity >= WORKSHOP_BUSY_RATIO)
  talleresOcupados.forEach(w => {
    items.push(insight(
      'info',
      `${w.name} cerca de su capacidad`,
      `${w.current}/${w.capacity} en uso. Próximo espacio libre: ${w.nextSlot}.`,
      'trazabilidad',
    ))
  })

  const noCertificados = workshops.filter(w => !w.certified)
  noCertificados.forEach(w => {
    items.push(insight(
      'advertencia',
      `${w.name} sin certificación vigente`,
      'No debería recibir contenedores de clientes con requisitos regulatorios hasta resolver esto.',
      'trazabilidad',
    ))
  })

  const severityOrder = { critico: 0, advertencia: 1, info: 2 }
  return items.sort((a, b) => severityOrder[a.severity] - severityOrder[b.severity])
}

function norm(text) {
  return text
    .toLowerCase()
    .normalize('NFD')
    .replace(/[̀-ͯ]/g, '')
    .trim()
}

function includesAny(text, keywords) {
  return keywords.some(k => text.includes(k))
}

function summarizeInsights(insights, max = 3) {
  if (insights.length === 0) {
    return 'No tengo nada urgente que reportarte en este momento — todo dentro de parámetros normales.'
  }
  const top = insights.slice(0, max)
  const lines = top.map(i => `• [${i.severity.toUpperCase()}] ${i.title} — ${i.message}`)
  const extra = insights.length > max ? `\n\nHay ${insights.length - max} más de menor prioridad si quieres el detalle completo.` : ''
  return lines.join('\n') + extra
}

// Lightweight intent matching over the mock data — no external model required.
// Swap this for a real LLM call (with these same data sources as context) when ready to go beyond canned intents.
export function answerQuery(rawText) {
  const text = norm(rawText || '')
  const insights = generateInsights()

  if (text.length === 0) {
    return 'Dime qué necesitas: flota, alertas, proyectos, inventario, talleres o un resumen general.'
  }

  if (includesAny(text, ['hola', 'buenos dias', 'buenas tardes', 'buenas noches', 'hey', 'que tal'])) {
    return `${getGreeting()}, Humberto. Ya revisé flota, proyectos, inventario y talleres. ${summarizeInsights(insights, 1)}`
  }

  if (includesAny(text, ['ayuda', 'que puedes hacer', 'help', 'opciones'])) {
    return 'Puedo darte el estado de la flota de contenedores, las alertas activas, el avance de proyectos, el inventario por bodega, o la ocupación de talleres. También te aviso proactivamente cuando detecto algo que necesita tu atención.'
  }

  if (includesAny(text, ['resumen', 'como vamos', 'estado general', 'todo bien', 'panorama'])) {
    const scKpis = kpis.supply_chain
    const kpiLine = scKpis.map(k => `${k.label}: ${k.value}`).join(' · ')
    return `Panorama general — ${kpiLine}.\n\n${summarizeInsights(insights)}`
  }

  if (includesAny(text, ['alerta', 'alertas', 'critico', 'urgente', 'problema'])) {
    return summarizeInsights(insights, 5)
  }

  if (includesAny(text, ['flota', 'contenedor', 'contenedores', 'ibc', 'tambor'])) {
    const byStatus = containers.reduce((acc, c) => {
      acc[c.status] = (acc[c.status] || 0) + 1
      return acc
    }, {})
    const line = Object.entries(byStatus).map(([status, count]) => `${status.replace('_', ' ')}: ${count}`).join(' · ')
    const perdidos = containers.filter(c => c.status === 'perdido').length
    const perdidosLine = perdidos > 0 ? ` Ojo: ${perdidos} sin ubicación confirmada.` : ''
    return `Flota total: ${containers.length} unidades. ${line}.${perdidosLine}`
  }

  if (includesAny(text, ['proyecto', 'proyectos', 'prj'])) {
    const enCurso = projects.filter(p => p.status === 'en_curso')
    const propuestas = projects.filter(p => p.status === 'propuesta')
    const riesgoAlto = projects.filter(p => p.risk === 'alto')
    return `${enCurso.length} proyecto(s) en curso, ${propuestas.length} en propuesta esperando decisión, ${riesgoAlto.length} en riesgo alto. ${riesgoAlto.length > 0 ? `Prioriza revisar: ${riesgoAlto.map(p => p.id).join(', ')}.` : 'Ninguno en riesgo alto por ahora.'}`
  }

  if (includesAny(text, ['stock', 'inventario', 'producto', 'bodega', 'marketplace'])) {
    const low = []
    Object.entries(WAREHOUSE_LABELS).forEach(([key, label]) => {
      products.forEach(p => {
        if (p.stock[key] < LOW_STOCK_THRESHOLD) low.push(`${p.name} en ${label} (${p.stock[key]})`)
      })
    })
    if (low.length === 0) return 'Inventario saludable en las tres bodegas, nada por debajo del mínimo.'
    return `Stock bajo mínimo:\n${low.map(l => `• ${l}`).join('\n')}`
  }

  if (includesAny(text, ['lavado', 'taller', 'talleres'])) {
    const line = workshops.map(w => `${w.name}: ${w.current}/${w.capacity}${w.certified ? '' : ' (sin certificación)'}`).join(' · ')
    return `Estado de talleres — ${line}.`
  }

  return 'Aún no tengo una respuesta preparada para eso. Puedo darte flota, alertas, proyectos, inventario o talleres — ¿por cuál empezamos?'
}
