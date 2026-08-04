export const roles = [
  { id: 'supply_chain', name: 'Supply Chain Director', avatar: 'SC', color: '#f97316', description: 'Visión global de operaciones, inventario y proyectos' },
  { id: 'plant_manager', name: 'Plant Manager', avatar: 'PM', color: '#3b82f6', description: 'Control de planta, lavado, reparaciones y calidad' },
  { id: 'logistics', name: 'Logistics Manager', avatar: 'LM', color: '#10b981', description: 'Rutas, despachos, entregas y trazabilidad en tránsito' },
  { id: 'packaging', name: 'Packaging Engineer', avatar: 'PE', color: '#8b5cf6', description: 'Especificaciones técnicas, validaciones y NC' },
  { id: 'ehs', name: 'EHS Manager', avatar: 'EH', color: '#f59e0b', description: 'Seguridad, medio ambiente y cumplimiento normativo' },
]

// Familias de activos reutilizables (taxonomía real de CHACONTAINER).
export const familias = [
  'Contenedor Colapsable',
  'Contenedor / Caja Rígida',
  'KLT',
  'Rack Metálico',
  'Tarima / Pallet',
  'Tapa',
  'Cinturón',
  'Charola',
  'Separador',
  'Inserto',
  'Dunnage',
]

export const containers = [
  { id: 'CHC-0421', type: 'Tarima/Pallet (CHA6548-50)', status: 'disponible', location: 'Bodega Estado de México', client: 'CMPC', washes: 12, lastEvent: '2026-05-15 08:32', lat: 19.6465, lng: -99.1968 },
  { id: 'CHC-0422', type: 'Tarima/Pallet (CHA6548-44)', status: 'en_lavado', location: 'Taller Querétaro', client: 'BASF', washes: 8, lastEvent: '2026-05-15 11:20', lat: 20.5888, lng: -100.3899 },
  { id: 'CHC-0423', type: 'KLT (CHA3032-25)', status: 'en_transito', location: 'Ruta Querétaro–CDMX', client: 'Codelco', washes: 15, lastEvent: '2026-05-15 09:45', lat: 20.10, lng: -99.90 },
  { id: 'CHC-0424', type: 'Contenedor Colapsable (CHA4845-42)', status: 'en_reparacion', location: 'Taller Puebla', client: 'ENAP', washes: 22, lastEvent: '2026-05-14 16:10', lat: 19.0414, lng: -98.2063 },
  { id: 'CHC-0425', type: 'Tarima/Pallet (CHA6548-50)', status: 'no_conforme', location: 'Bodega Puebla', client: 'CMPC', washes: 31, lastEvent: '2026-05-13 14:55', lat: 19.0414, lng: -98.2063 },
  { id: 'CHC-0426', type: 'KLT (CHA3032-34)', status: 'asignado', location: 'Planta Cliente Puebla', client: 'Arauco', washes: 5, lastEvent: '2026-05-15 07:30', lat: 19.03, lng: -98.19 },
  { id: 'CHC-0427', type: 'Tarima/Pallet (CHA6548-34)', status: 'disponible', location: 'Bodega Estado de México', client: null, washes: 3, lastEvent: '2026-05-14 12:00', lat: 19.6465, lng: -99.1968 },
  { id: 'CHC-0428', type: 'Tarima/Pallet (CHA6548-50)', status: 'perdido', location: 'Desconocido', client: 'BASF', washes: 18, lastEvent: '2026-05-10 09:00', lat: null, lng: null },
  { id: 'CHC-0429', type: 'Contenedor Colapsable (CHA4845-25)', status: 'disponible', location: 'Bodega Querétaro', client: null, washes: 7, lastEvent: '2026-05-15 06:15', lat: 20.5888, lng: -100.3899 },
  { id: 'CHC-0430', type: 'KLT (CHA3032-25)', status: 'en_lavado', location: 'Taller Villa de Reyes', client: 'Arauco', washes: 11, lastEvent: '2026-05-15 10:00', lat: 21.8078, lng: -100.9271 },
  { id: 'CHC-0431', type: 'Rack Metálico (CHA64.548-34)', status: 'asignado', location: 'Planta Cliente Toluca', client: 'CMPC', washes: 6, lastEvent: '2026-05-15 08:00', lat: 19.2826, lng: -99.6557 },
  { id: 'CHC-0432', type: 'Contenedor Colapsable (CHA4845-34)', status: 'en_transito', location: 'Ruta San Luis Potosí–Querétaro', client: 'Codelco', washes: 9, lastEvent: '2026-05-15 05:30', lat: 22.1565, lng: -100.9855 },
  { id: 'CHC-0433', type: 'Contenedor / Caja Rígida (CHA4856-34)', status: 'disponible', location: 'Bodega Estado de México', client: null, washes: 14, lastEvent: '2026-05-14 15:20', lat: 19.6465, lng: -99.1968 },
  { id: 'CHC-0434', type: 'Charola (CHA7048-50)', status: 'en_reparacion', location: 'Taller Querétaro', client: 'ENAP', washes: 27, lastEvent: '2026-05-14 11:45', lat: 20.5888, lng: -100.3899 },
  { id: 'CHC-0435', type: 'KLT (CHA3032-34)', status: 'disponible', location: 'Bodega Puebla', client: null, washes: 4, lastEvent: '2026-05-13 09:30', lat: 19.0414, lng: -98.2063 },
  { id: 'CHC-0436', type: 'Rack Metálico (CHA64.548-25)', status: 'asignado', location: 'Planta Cliente Guanajuato', client: 'ENAP', washes: 19, lastEvent: '2026-05-15 07:00', lat: 21.0190, lng: -101.2574 },
  { id: 'CHC-0437', type: 'Contenedor Colapsable (CHA4845-42)', status: 'no_conforme', location: 'Bodega Estado de México', client: 'BASF', washes: 33, lastEvent: '2026-05-12 14:00', lat: 19.6465, lng: -99.1968 },
  { id: 'CHC-0438', type: 'Separador (CHA7848-34)', status: 'en_lavado', location: 'Taller Puebla', client: 'Codelco', washes: 10, lastEvent: '2026-05-15 09:15', lat: 19.0414, lng: -98.2063 },
  { id: 'CHC-0439', type: 'Tarima/Pallet (CHA6548-25)', status: 'disponible', location: 'Bodega Querétaro', client: null, washes: 2, lastEvent: '2026-05-15 06:45', lat: 20.5888, lng: -100.3899 },
  { id: 'CHC-0440', type: 'Contenedor / Caja Rígida (CHA4857-34)', status: 'en_transito', location: 'Ruta Puebla–Villa de Reyes', client: 'Arauco', washes: 8, lastEvent: '2026-05-15 08:30', lat: 20.42, lng: -99.57 },
]

export const containerEvents = {
  'CHC-0421': [
    { time: '2026-05-15 08:32', event: 'Ingreso Bodega', detail: 'Bodega Estado de México - Zona A3', user: 'Sistema QR' },
    { time: '2026-05-14 17:15', event: 'Lavado Completo', detail: 'Ciclo 12 finalizado - Aprobado QC', user: 'Taller Qro' },
    { time: '2026-05-14 09:00', event: 'Ingreso Lavado', detail: 'Taller Querétaro Centro', user: 'Pedro L.' },
    { time: '2026-05-13 14:30', event: 'Devolución Cliente', detail: 'CMPC Planta Toluca → Taller', user: 'Transportes XY' },
  ],
  'CHC-0422': [
    { time: '2026-05-15 11:20', event: 'En Lavado', detail: 'Taller Querétaro - Iniciado ciclo 9', user: 'Operador M.' },
    { time: '2026-05-15 09:45', event: 'Ingreso Taller', detail: 'Devolución BASF Querétaro', user: 'Camión XYZ-123' },
    { time: '2026-05-10 08:00', event: 'Despacho', detail: 'Enviado a BASF Querétaro', user: 'Sistema' },
  ],
  'CHC-0428': [
    { time: '2026-05-10 09:00', event: 'Último Escaneo', detail: 'BASF Planta Estado de México - Ingreso', user: 'Guardia BASF' },
    { time: '2026-05-08 14:00', event: 'Despacho', detail: 'Enviado desde Bodega Estado de México', user: 'Sistema' },
    { time: '2026-05-05 11:30', event: 'Lavado Completo', detail: 'Ciclo 18 finalizado', user: 'Taller Qro' },
  ],
}

export const products = [
  { id: 'P001', name: 'KLT Chico', sku: 'CHA3032-25', familia: 'KLT', category: 'KLT', price: 18000, stock: { bodega_edomex: 120, bodega_queretaro: 85, bodega_puebla: 40 }, unit: 'unidad', minOrder: 10, co2Saved: 1.2 },
  { id: 'P002', name: 'Contenedor Colapsable Mediano', sku: 'CHA4845-34', familia: 'Contenedor Colapsable', category: 'Contenedor Colapsable', price: 195000, stock: { bodega_edomex: 15, bodega_queretaro: 28, bodega_puebla: 3 }, unit: 'unidad', minOrder: 3, co2Saved: 8.2 },
  { id: 'P003', name: 'Contenedor / Caja Rígida', sku: 'CHA4856-34', familia: 'Contenedor / Caja Rígida', category: 'Contenedor / Caja Rígida', price: 210000, stock: { bodega_edomex: 42, bodega_queretaro: 18, bodega_puebla: 7 }, unit: 'unidad', minOrder: 5, co2Saved: 9.4 },
  { id: 'P004', name: 'Rack Metálico', sku: 'CHA64.548-34', familia: 'Rack Metálico', category: 'Rack Metálico', price: 285000, stock: { bodega_edomex: 8, bodega_queretaro: 12, bodega_puebla: 2 }, unit: 'unidad', minOrder: 2, co2Saved: 13.1 },
  { id: 'P005', name: 'Tarima / Pallet', sku: 'CHA6548-50', familia: 'Tarima / Pallet', category: 'Tarima / Pallet', price: 45000, stock: { bodega_edomex: 200, bodega_queretaro: 150, bodega_puebla: 80 }, unit: 'unidad', minOrder: 20, co2Saved: 2.1 },
  { id: 'P006', name: 'Charola', sku: 'CHA7048-50', familia: 'Charola', category: 'Charola', price: 52000, stock: { bodega_edomex: 60, bodega_queretaro: 45, bodega_puebla: 20 }, unit: 'unidad', minOrder: 10, co2Saved: 1.8 },
  { id: 'P007', name: 'Separador', sku: 'CHA7848-34', familia: 'Separador', category: 'Separador', price: 12000, stock: { bodega_edomex: 300, bodega_queretaro: 210, bodega_puebla: 95 }, unit: 'unidad', minOrder: 20, co2Saved: 0.5 },
  { id: 'P008', name: 'Servicio de Lavado', sku: 'SRV-LAV', familia: null, category: 'Servicios', price: 28000, stock: { bodega_edomex: 999, bodega_queretaro: 999, bodega_puebla: 999 }, unit: 'unidad', minOrder: 1, co2Saved: 0 },
  { id: 'P009', name: 'Servicio de Reparación', sku: 'SRV-REP', familia: null, category: 'Servicios', price: 65000, stock: { bodega_edomex: 999, bodega_queretaro: 999, bodega_puebla: 999 }, unit: 'unidad', minOrder: 1, co2Saved: 0 },
]

export const projects = [
  { id: 'PRJ-001', name: 'Migración Embalaje CMPC Planta Toluca', client: 'CMPC', manager: 'Andrea Vásquez', status: 'en_curso', phase: 3, totalPhases: 5, startDate: '2026-03-01', endDate: '2026-07-30', budget: 45000000, spent: 28500000, containers: 180, risk: 'medio' },
  { id: 'PRJ-002', name: 'Implementación Contenedores BASF Querétaro', client: 'BASF', manager: 'Carlos Muñoz', status: 'en_curso', phase: 2, totalPhases: 4, startDate: '2026-04-15', endDate: '2026-09-15', budget: 32000000, spent: 12800000, containers: 120, risk: 'bajo' },
  { id: 'PRJ-003', name: 'Piloto Tarimas Codelco Puebla', client: 'Codelco', manager: 'María Torres', status: 'propuesta', phase: 1, totalPhases: 3, startDate: '2026-06-01', endDate: '2026-10-31', budget: 18000000, spent: 0, containers: 300, risk: 'alto' },
  { id: 'PRJ-004', name: 'Renovación Flota ENAP Guanajuato', client: 'ENAP', manager: 'Roberto Silva', status: 'cerrado', phase: 5, totalPhases: 5, startDate: '2025-09-01', endDate: '2026-03-31', budget: 28000000, spent: 27200000, containers: 95, risk: 'bajo' },
  { id: 'PRJ-005', name: 'Expansión Arauco Villa de Reyes', client: 'Arauco', manager: 'Andrea Vásquez', status: 'en_curso', phase: 1, totalPhases: 6, startDate: '2026-05-01', endDate: '2026-12-31', budget: 55000000, spent: 4200000, containers: 250, risk: 'medio' },
  { id: 'PRJ-006', name: 'Evaluación Contenedores BASF Estado de México', client: 'BASF', manager: 'Carlos Muñoz', status: 'propuesta', phase: 1, totalPhases: 4, startDate: '2026-07-01', endDate: '2027-01-31', budget: 22000000, spent: 0, containers: 80, risk: 'bajo' },
]

export const alerts = [
  { id: 'A001', type: 'critico', message: 'CHC-0428 sin escaneo hace 5 días - posible pérdida', module: 'trazabilidad', time: '10:32' },
  { id: 'A002', type: 'advertencia', message: 'Stock Tarima/Pallet Bodega Querétaro bajo mínimo (7 un)', module: 'marketplace', time: '09:15' },
  { id: 'A003', type: 'info', message: 'Pedido #ORD-2847 despachado desde Bodega Estado de México', module: 'marketplace', time: '08:50' },
  { id: 'A004', type: 'advertencia', message: 'PRJ-003 Codelco: aprobación pendiente hace 3 días', module: 'proyectos', time: '08:20' },
  { id: 'A005', type: 'critico', message: '3 contenedores con vencimiento de certificación esta semana', module: 'trazabilidad', time: 'ayer' },
]

export const kpis = {
  supply_chain: [
    { label: 'Contenedores Activos', value: '1,247', delta: '+3.2%', trend: 'up', icon: 'Package' },
    { label: 'Rotación Mensual', value: '89.4%', delta: '+1.1%', trend: 'up', icon: 'RefreshCw' },
    { label: 'Pedidos Pendientes', value: '34', delta: '-8', trend: 'up', icon: 'ShoppingCart' },
    { label: 'CO₂ Evitado (ton)', value: '128.4', delta: '+12.3', trend: 'up', icon: 'Leaf' },
    { label: 'Tasa Pérdidas', value: '0.08%', delta: '-0.02%', trend: 'up', icon: 'AlertTriangle' },
    { label: 'Proyectos Activos', value: '3', delta: '0', trend: 'neutral', icon: 'FolderOpen' },
  ],
  plant_manager: [
    { label: 'Disponibles en Planta', value: '84', delta: '+6', trend: 'up', icon: 'Package' },
    { label: 'En Lavado', value: '12', delta: '-2', trend: 'up', icon: 'Droplets' },
    { label: 'En Reparación', value: '5', delta: '+1', trend: 'down', icon: 'Wrench' },
    { label: 'No Conformes', value: '2', delta: '+2', trend: 'down', icon: 'XCircle' },
    { label: 'Ciclo Promedio (días)', value: '4.2', delta: '-0.3', trend: 'up', icon: 'Clock' },
    { label: 'Eficiencia Lavado', value: '96.8%', delta: '+0.5%', trend: 'up', icon: 'CheckCircle' },
  ],
  logistics: [
    { label: 'En Tránsito', value: '47', delta: '+5', trend: 'neutral', icon: 'Truck' },
    { label: 'Entregas Hoy', value: '12', delta: '0', trend: 'neutral', icon: 'MapPin' },
    { label: 'Pedidos Procesados', value: '189', delta: '+23', trend: 'up', icon: 'ShoppingCart' },
    { label: 'Tiempo Entrega Prom.', value: '1.8 días', delta: '-0.2', trend: 'up', icon: 'Clock' },
    { label: 'Incidencias Ruta', value: '1', delta: '-2', trend: 'up', icon: 'AlertTriangle' },
    { label: 'Cobertura Zonas', value: '7 regiones', delta: '0', trend: 'neutral', icon: 'Map' },
  ],
  packaging: [
    { label: 'Especificaciones Activas', value: '23', delta: '+2', trend: 'up', icon: 'FileText' },
    { label: 'Validaciones Pendientes', value: '4', delta: '-1', trend: 'up', icon: 'CheckSquare' },
    { label: 'Tipos de Contenedor', value: '8', delta: '0', trend: 'neutral', icon: 'Layers' },
    { label: 'Proyectos Asignados', value: '2', delta: '0', trend: 'neutral', icon: 'FolderOpen' },
    { label: 'NC Reportadas (mes)', value: '3', delta: '-2', trend: 'up', icon: 'AlertCircle' },
    { label: 'Ensayos Programados', value: '6', delta: '+3', trend: 'neutral', icon: 'FlaskConical' },
  ],
  ehs: [
    { label: 'Días sin Incidentes', value: '47', delta: '+1', trend: 'up', icon: 'Shield' },
    { label: 'Contenedores NC', value: '2', delta: '+2', trend: 'down', icon: 'XCircle' },
    { label: 'Certificaciones Venc.', value: '3', delta: '+3', trend: 'down', icon: 'AlertTriangle' },
    { label: 'CO₂ Evitado (ton)', value: '128.4', delta: '+12.3', trend: 'up', icon: 'Leaf' },
    { label: 'Auditorías Pendientes', value: '1', delta: '0', trend: 'neutral', icon: 'ClipboardCheck' },
    { label: 'Tasa Reciclaje', value: '94.2%', delta: '+1.8%', trend: 'up', icon: 'Recycle' },
  ],
}

export const recentActivity = [
  { time: '11:42', event: 'QR Escaneado', detail: 'CHC-0421 → Bodega Estado de México - Ingreso', user: 'Carlos M.' },
  { time: '11:38', event: 'Pedido Creado', detail: 'ORD-2851: 20x Tarima/Pallet → CMPC Toluca', user: 'Sistema' },
  { time: '11:25', event: 'Lavado Completo', detail: 'CHC-0419 → Disponible (ciclo 9)', user: 'Taller Qro' },
  { time: '11:10', event: 'NC Registrada', detail: 'CHC-0425 → Válvula dañada - Requiere reparación', user: 'Pedro L.' },
  { time: '10:55', event: 'Despacho', detail: 'ORD-2847: 15x Contenedor Colapsable → ENAP Guanajuato', user: 'Sistema' },
  { time: '10:32', event: 'Alerta Crítica', detail: 'CHC-0428 sin escaneo hace 5 días', user: 'Sistema' },
]

export const workshops = [
  { id: 'TLR-001', name: 'Taller Querétaro', capacity: 80, current: 12, certified: true, nextSlot: '2026-05-16 14:00' },
  { id: 'TLR-002', name: 'Taller Puebla', capacity: 60, current: 5, certified: true, nextSlot: '2026-05-16 09:00' },
  { id: 'TLR-003', name: 'Taller Villa de Reyes', capacity: 45, current: 8, certified: false, nextSlot: '2026-05-17 10:00' },
  { id: 'TLR-004', name: 'Taller Guanajuato', capacity: 30, current: 3, certified: true, nextSlot: '2026-05-16 16:00' },
]

export const statusConfig = {
  disponible: { label: 'Disponible', color: '#10b981', bg: 'bg-emerald-900/40', text: 'text-emerald-400', border: 'border-emerald-700' },
  en_lavado: { label: 'En Lavado', color: '#3b82f6', bg: 'bg-blue-900/40', text: 'text-blue-400', border: 'border-blue-700' },
  en_transito: { label: 'En Tránsito', color: '#f59e0b', bg: 'bg-amber-900/40', text: 'text-amber-400', border: 'border-amber-700' },
  en_reparacion: { label: 'En Reparación', color: '#8b5cf6', bg: 'bg-violet-900/40', text: 'text-violet-400', border: 'border-violet-700' },
  asignado: { label: 'Asignado', color: '#06b6d4', bg: 'bg-cyan-900/40', text: 'text-cyan-400', border: 'border-cyan-700' },
  no_conforme: { label: 'No Conforme', color: '#ef4444', bg: 'bg-red-900/40', text: 'text-red-400', border: 'border-red-700' },
  perdido: { label: 'Perdido', color: '#6b7280', bg: 'bg-gray-900/40', text: 'text-gray-400', border: 'border-gray-600' },
}

export const chartDataContainers = [
  { name: 'Disponible', value: 8, fill: '#10b981' },
  { name: 'En Lavado', value: 3, fill: '#3b82f6' },
  { name: 'En Tránsito', value: 3, fill: '#f59e0b' },
  { name: 'Asignado', value: 3, fill: '#06b6d4' },
  { name: 'En Reparación', value: 2, fill: '#8b5cf6' },
  { name: 'No Conforme', value: 2, fill: '#ef4444' },
  { name: 'Perdido', value: 1, fill: '#6b7280' },
]

export const monthlyData = [
  { mes: 'Nov', ciclos: 312, co2: 98.4, pedidos: 124 },
  { mes: 'Dic', ciclos: 289, co2: 91.2, pedidos: 108 },
  { mes: 'Ene', ciclos: 334, co2: 105.4, pedidos: 132 },
  { mes: 'Feb', ciclos: 358, co2: 113.0, pedidos: 145 },
  { mes: 'Mar', ciclos: 401, co2: 126.6, pedidos: 162 },
  { mes: 'Abr', ciclos: 378, co2: 119.3, pedidos: 151 },
  { mes: 'May', ciclos: 406, co2: 128.4, pedidos: 189 },
]
