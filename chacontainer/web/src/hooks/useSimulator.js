import { useEffect, useRef } from 'react'
import { useApp } from '../context/AppContext'

const scanEvents = [
  { containerId: 'CHC-0421', location: 'Bodega Estado de México', user: 'Carlos M.' },
  { containerId: 'CHC-0422', location: 'Taller Querétaro', user: 'Operador R.' },
  { containerId: 'CHC-0430', location: 'Taller Villa de Reyes', user: 'María T.' },
  { containerId: 'CHC-0426', location: 'Planta Cliente Puebla', user: 'Guardia A.' },
  { containerId: 'CHC-0433', location: 'Bodega Estado de México', user: 'Sistema QR' },
  { containerId: 'CHC-0439', location: 'Bodega Querétaro', user: 'Juan P.' },
]

const orderEvents = [
  { order: 'ORD-2852', detail: '10x Contenedor Colapsable → Valeo México Querétaro', user: 'Sistema' },
  { order: 'ORD-2853', detail: '30x KLT → Denso Puebla', user: 'Sistema' },
  { order: 'ORD-2854', detail: '5x Tarima/Pallet → Continental Guanajuato', user: 'Sistema' },
]

const washEvents = [
  { containerId: 'CHC-0419', cycle: 9, user: 'Taller Qro' },
  { containerId: 'CHC-0438', cycle: 11, user: 'Taller Puebla' },
  { containerId: 'CHC-0430', cycle: 12, user: 'Taller VdR' },
]

function getRandomInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min
}

function now() {
  const d = new Date()
  return `${String(d.getHours()).padStart(2, '0')}:${String(d.getMinutes()).padStart(2, '0')}`
}

export function useSimulator() {
  const { addActivity, addAlert } = useApp()
  const timerRef = useRef(null)
  const tickRef = useRef(0)

  useEffect(() => {
    function tick() {
      tickRef.current += 1
      const t = tickRef.current
      const time = now()

      if (t % 1 === 0) {
        // QR scan event
        const scan = scanEvents[getRandomInt(0, scanEvents.length - 1)]
        addActivity({
          time,
          event: 'QR Escaneado',
          detail: `${scan.containerId} → ${scan.location}`,
          user: scan.user,
        })
      }

      if (t % 2 === 0) {
        // Order or wash event
        if (Math.random() > 0.5) {
          const order = orderEvents[getRandomInt(0, orderEvents.length - 1)]
          addActivity({
            time,
            event: 'Pedido Creado',
            detail: `${order.order}: ${order.detail}`,
            user: order.user,
          })
        } else {
          const wash = washEvents[getRandomInt(0, washEvents.length - 1)]
          addActivity({
            time,
            event: 'Lavado Completo',
            detail: `${wash.containerId} → Disponible (ciclo ${wash.cycle})`,
            user: wash.user,
          })
        }
      }

      if (t % 4 === 0 && Math.random() > 0.6) {
        // Random alert
        const alertTypes = [
          { type: 'info', message: `Despacho completado: ${getRandomInt(5, 20)} contenedores → Cliente`, module: 'marketplace' },
          { type: 'advertencia', message: `Stock bajo en Bodega Puebla: Contenedor Colapsable (${getRandomInt(2, 8)} un)`, module: 'marketplace' },
          { type: 'info', message: `Ciclo de lavado #${getRandomInt(8, 15)} completado en Taller Querétaro`, module: 'trazabilidad' },
        ]
        const alert = alertTypes[getRandomInt(0, alertTypes.length - 1)]
        addAlert({ id: `A${Date.now()}`, ...alert, time })
      }

      const delay = getRandomInt(8000, 15000)
      timerRef.current = setTimeout(tick, delay)
    }

    timerRef.current = setTimeout(tick, getRandomInt(8000, 12000))
    return () => {
      if (timerRef.current) clearTimeout(timerRef.current)
    }
  }, [addActivity, addAlert])
}
