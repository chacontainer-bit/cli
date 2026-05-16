export const roles = [
  { id: 'supply_chain', name: 'Supply Chain Director', avatar: 'SC', color: '#f97316', description: 'Visión global de operaciones, inventario y proyectos' },
  { id: 'plant_manager', name: 'Plant Manager', avatar: 'PM', color: '#3b82f6', description: 'Control de planta, lavado, reparaciones y calidad' },
  { id: 'logistics', name: 'Logistics Manager', avatar: 'LM', color: '#10b981', description: 'Rutas, despachos, entregas y trazabilidad en tránsito' },
  { id: 'packaging', name: 'Packaging Engineer', avatar: 'PE', color: '#8b5cf6', description: 'Especificaciones técnicas, validaciones y NC' },
  { id: 'ehs', name: 'EHS Manager', avatar: 'EH', color: '#f59e0b', description: 'Seguridad, medio ambiente y cumplimiento normativo' },
]

export const containers = [
  { id: 'CHC-0421', type: 'IBC 1000L', status: 'disponible', location: 'Bodega Norte', client: 'CMPC', washes: 12, lastEvent: '2026-05-15 08:32', lat: -33.45, lng: -70.65 },
  { id: 'CHC-0422', type: 'IBC 1000L', status: 'en_lavado', location: 'Taller Santiago', client: 'BASF', washes: 8, lastEvent: '2026-05-15 11:20', lat: -33.47, lng: -70.67 },
  { id: 'CHC-0423', type: 'Tambor 200L', status: 'en_transito', location: 'Ruta Valparaíso', client: 'Codelco', washes: 15, lastEvent: '2026-05-15 09:45', lat: -33.02, lng: -71.45 },
  { id: 'CHC-0424', type: 'IBC 600L', status: 'en_reparacion', location: 'Taller Rancagua', client: 'ENAP', washes: 22, lastEvent: '2026-05-14 16:10', lat: -34.17, lng: -70.74 },
  { id: 'CHC-0425', type: 'IBC 1000L', status: 'no_conforme', location: 'Bodega Sur', client: 'CMPC', washes: 31, lastEvent: '2026-05-13 14:55', lat: -33.50, lng: -70.62 },
  { id: 'CHC-0426', type: 'Tambor 200L', status: 'asignado', location: 'Planta Concepción', client: 'Arauco', washes: 5, lastEvent: '2026-05-15 07:30', lat: -36.82, lng: -73.05 },
  { id: 'CHC-0427', type: 'IBC 1000L', status: 'disponible', location: 'Bodega Norte', client: null, washes: 3, lastEvent: '2026-05-14 12:00', lat: -33.45, lng: -70.64 },
  { id: 'CHC-0428', type: 'IBC 1000L', status: 'perdido', location: 'Desconocido', client: 'BASF', washes: 18, lastEvent: '2026-05-10 09:00', lat: null, lng: null },
  { id: 'CHC-0429', type: 'IBC 600L', status: 'disponible', location: 'Bodega Valparaíso', client: null, washes: 7, lastEvent: '2026-05-15 06:15', lat: -33.04, lng: -71.46 },
  { id: 'CHC-0430', type: 'Tambor 200L', status: 'en_lavado', location: 'Taller Concepción', client: 'Arauco', washes: 11, lastEvent: '2026-05-15 10:00', lat: -36.83, lng: -73.06 },
  { id: 'CHC-0431', type: 'IBC 1000L', status: 'asignado', location: 'Planta CMPC Nacimiento', client: 'CMPC', washes: 6, lastEvent: '2026-05-15 08:00', lat: -37.51, lng: -72.68 },
  { id: 'CHC-0432', type: 'IBC 600L', status: 'en_transito', location: 'Ruta Antofagasta', client: 'Codelco', washes: 9, lastEvent: '2026-05-15 05:30', lat: -23.65, lng: -70.40 },
  { id: 'CHC-0433', type: 'Tambor 220L', status: 'disponible', location: 'Bodega Norte', client: null, washes: 14, lastEvent: '2026-05-14 15:20', lat: -33.45, lng: -70.66 },
  { id: 'CHC-0434', type: 'IBC 1000L', status: 'en_reparacion', location: 'Taller Santiago', client: 'ENAP', washes: 27, lastEvent: '2026-05-14 11:45', lat: -33.47, lng: -70.67 },
  { id: 'CHC-0435', type: 'Tambor 200L', status: 'disponible', location: 'Bodega Sur', client: null, washes: 4, lastEvent: '2026-05-13 09:30', lat: -33.50, lng: -70.62 },
  { id: 'CHC-0436', type: 'IBC 1000L', status: 'asignado', location: 'Refinería ENAP', client: 'ENAP', washes: 19, lastEvent: '2026-05-15 07:00', lat: -32.85, lng: -71.50 },
  { id: 'CHC-0437', type: 'IBC 600L', status: 'no_conforme', location: 'Bodega Norte', client: 'BASF', washes: 33, lastEvent: '2026-05-12 14:00', lat: -33.45, lng: -70.65 },
  { id: 'CHC-0438', type: 'Tambor 200L', status: 'en_lavado', location: 'Taller Rancagua', client: 'Codelco', washes: 10, lastEvent: '2026-05-15 09:15', lat: -34.17, lng: -70.74 },
  { id: 'CHC-0439', type: 'IBC 1000L', status: 'disponible', location: 'Bodega Valparaíso', client: null, washes: 2, lastEvent: '2026-05-15 06:45', lat: -33.04, lng: -71.46 },
  { id: 'CHC-0440', type: 'Tambor 220L', status: 'en_transito', location: 'Ruta Santiago-Concepción', client: 'Arauco', washes: 8, lastEvent: '2026-05-15 08:30', lat: -35.10, lng: -71.60 },
]

export const containerEvents = {
  'CHC-0421': [
    { time: '2026-05-15 08:32', event: 'Ingreso Bodega', detail: 'Bodega Norte - Zona A3', user: 'Sistema QR' },
    { time: '2026-05-14 17:15', event: 'Lavado Completo', detail: 'Ciclo 12 finalizado - Aprobado QC', user: 'Taller Stgo' },
    { time: '2026-05-14 09:00', event: 'Ingreso Lavado', detail: 'Taller Santiago Centro', user: 'Pedro L.' },
    { time: '2026-05-13 14:30', event: 'Devolución Cliente', detail: 'CMPC Planta Nacimiento → Taller', user: 'Transportes XY' },
  ],
  'CHC-0422': [
    { time: '2026-05-15 11:20', event: 'En Lavado', detail: 'Taller Santiago - Iniciado ciclo 9', user: 'Operador M.' },
    { time: '2026-05-15 09:45', event: 'Ingreso Taller', detail: 'Devolución BASF Santiago', user: 'Camión XYZ-123' },
    { time: '2026-05-10 08:00', event: 'Despacho', detail: 'Enviado a BASF Santiago', user: 'Sistema' },
  ],
  'CHC-0428': [
    { time: '2026-05-10 09:00', event: 'Último Escaneo', detail: 'BASF Planta Norte - Ingreso', user: 'Guardia BASF' },
    { time: '2026-05-08 14:00', event: 'Despacho', detail: 'Enviado desde Bodega Norte', user: 'Sistema' },
    { time: '2026-05-05 11:30', event: 'Lavado Completo', detail: 'Ciclo 18 finalizado', user: 'Taller Stgo' },
  ],
}

export const products = [
  { id: 'P001', name: 'IBC 1000L Grado Alimenticio', sku: 'IBC-1000-FA', category: 'IBC', price: 285000, stock: { bodega_norte: 42, bodega_sur: 18, bodega_valpo: 7 }, unit: 'unidad', minOrder: 5, co2Saved: 12.4 },
  { id: 'P002', name: 'IBC 600L Industrial', sku: 'IBC-600-IND', category: 'IBC', price: 195000, stock: { bodega_norte: 15, bodega_sur: 28, bodega_valpo: 3 }, unit: 'unidad', minOrder: 3, co2Saved: 8.2 },
  { id: 'P003', name: 'Tambor 200L HDPE', sku: 'TAM-200-HDPE', category: 'Tambores', price: 45000, stock: { bodega_norte: 120, bodega_sur: 85, bodega_valpo: 40 }, unit: 'unidad', minOrder: 10, co2Saved: 2.1 },
  { id: 'P004', name: 'Tambor 220L Metálico', sku: 'TAM-220-MET', category: 'Tambores', price: 52000, stock: { bodega_norte: 60, bodega_sur: 45, bodega_valpo: 20 }, unit: 'unidad', minOrder: 10, co2Saved: 1.8 },
  { id: 'P005', name: 'Servicio Lavado IBC', sku: 'SRV-LAV-IBC', category: 'Servicios', price: 28000, stock: { bodega_norte: 999, bodega_sur: 999, bodega_valpo: 999 }, unit: 'unidad', minOrder: 1, co2Saved: 0 },
  { id: 'P006', name: 'Servicio Reparación IBC', sku: 'SRV-REP-IBC', category: 'Servicios', price: 65000, stock: { bodega_norte: 999, bodega_sur: 999, bodega_valpo: 999 }, unit: 'unidad', minOrder: 1, co2Saved: 0 },
  { id: 'P007', name: 'IBC 1000L Grado Químico', sku: 'IBC-1000-QM', category: 'IBC', price: 310000, stock: { bodega_norte: 8, bodega_sur: 12, bodega_valpo: 2 }, unit: 'unidad', minOrder: 2, co2Saved: 13.1 },
  { id: 'P008', name: 'Pallet Reutilizable', sku: 'PAL-RUT-STD', category: 'Pallets', price: 18000, stock: { bodega_norte: 200, bodega_sur: 150, bodega_valpo: 80 }, unit: 'unidad', minOrder: 20, co2Saved: 0.5 },
]

export const projects = [
  { id: 'PRJ-001', name: 'Migración Embalaje CMPC Planta Nacimiento', client: 'CMPC', manager: 'Andrea Vásquez', status: 'en_curso', phase: 3, totalPhases: 5, startDate: '2026-03-01', endDate: '2026-07-30', budget: 45000000, spent: 28500000, containers: 180, risk: 'medio' },
  { id: 'PRJ-002', name: 'Implementación IBC BASF Región Metropolitana', client: 'BASF', manager: 'Carlos Muñoz', status: 'en_curso', phase: 2, totalPhases: 4, startDate: '2026-04-15', endDate: '2026-09-15', budget: 32000000, spent: 12800000, containers: 120, risk: 'bajo' },
  { id: 'PRJ-003', name: 'Piloto Tambores Codelco División Andina', client: 'Codelco', manager: 'María Torres', status: 'propuesta', phase: 1, totalPhases: 3, startDate: '2026-06-01', endDate: '2026-10-31', budget: 18000000, spent: 0, containers: 300, risk: 'alto' },
  { id: 'PRJ-004', name: 'Renovación Flota IBC ENAP Refinería', client: 'ENAP', manager: 'Roberto Silva', status: 'cerrado', phase: 5, totalPhases: 5, startDate: '2025-09-01', endDate: '2026-03-31', budget: 28000000, spent: 27200000, containers: 95, risk: 'bajo' },
  { id: 'PRJ-005', name: 'Expansión Arauco Planta Horcones', client: 'Arauco', manager: 'Andrea Vásquez', status: 'en_curso', phase: 1, totalPhases: 6, startDate: '2026-05-01', endDate: '2026-12-31', budget: 55000000, spent: 4200000, containers: 250, risk: 'medio' },
  { id: 'PRJ-006', name: 'Evaluación Contenedores Química RM', client: 'BASF', manager: 'Carlos Muñoz', status: 'propuesta', phase: 1, totalPhases: 4, startDate: '2026-07-01', endDate: '2027-01-31', budget: 22000000, spent: 0, containers: 80, risk: 'bajo' },
]

export const alerts = [
  { id: 'A001', type: 'critico', message: 'CHC-0428 sin escaneo hace 5 días - posible pérdida', module: 'trazabilidad', time: '10:32' },
  { id: 'A002', type: 'advertencia', message: 'Stock IBC 1000L Bodega Valparaíso bajo mínimo (7 un)', module: 'marketplace', time: '09:15' },
  { id: 'A003', type: 'info', message: 'Pedido #ORD-2847 despachado desde Bodega Norte', module: 'marketplace', time: '08:50' },
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
  { time: '11:42', event: 'QR Escaneado', detail: 'CHC-0421 → Bodega Norte - Ingreso', user: 'Carlos M.' },
  { time: '11:38', event: 'Pedido Creado', detail: 'ORD-2851: 20x IBC 1000L → CMPC Nacimiento', user: 'Sistema' },
  { time: '11:25', event: 'Lavado Completo', detail: 'CHC-0419 → Disponible (ciclo 9)', user: 'Taller Stgo' },
  { time: '11:10', event: 'NC Registrada', detail: 'CHC-0425 → Válvula dañada - Requiere reparación', user: 'Pedro L.' },
  { time: '10:55', event: 'Despacho', detail: 'ORD-2847: 15x Tambor 200L → ENAP Refinería', user: 'Sistema' },
  { time: '10:32', event: 'Alerta Crítica', detail: 'CHC-0428 sin escaneo hace 5 días', user: 'Sistema' },
]

export const workshops = [
  { id: 'TLR-001', name: 'Taller Santiago Centro', capacity: 80, current: 12, certified: true, nextSlot: '2026-05-16 14:00' },
  { id: 'TLR-002', name: 'Taller Rancagua', capacity: 60, current: 5, certified: true, nextSlot: '2026-05-16 09:00' },
  { id: 'TLR-003', name: 'Taller Concepción', capacity: 45, current: 8, certified: false, nextSlot: '2026-05-17 10:00' },
  { id: 'TLR-004', name: 'Taller Antofagasta', capacity: 30, current: 3, certified: true, nextSlot: '2026-05-16 16:00' },
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
