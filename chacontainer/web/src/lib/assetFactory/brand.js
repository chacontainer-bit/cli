// Design tokens de marca CHACONTAINER, tomados de tailwind.config.js
// (chacontainer/web) para que las piezas generadas usen el mismo sistema
// visual que el resto de la app.
export const BRAND = {
  name: 'CHACONTAINER',
  tagline: 'Contenedores reutilizables',
  color: '#f97316',
  colorDark: '#ea6c0a',
  surface: '#111827',
  surfaceDark: '#0a0f1a',
  surfaceCard: '#1a2235',
  border: '#1f2937',
  textMuted: '#9ca3af',
  textFaint: '#6b7280',
  textLight: '#f9fafb',
  fontFamily: "'Inter', system-ui, sans-serif",
}

// Tamaños de página en px (96dpi) usados para renderizar los nodos que luego
// se capturan 1:1 hacia el PDF, evitando distorsión de aspecto.
export const PAGE_A4 = { width: 794, height: 1123 }
export const SLIDE_16_9 = { width: 1280, height: 720 }
