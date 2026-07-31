const DIACRITICS_RE = new RegExp('[\\u0300-\\u036f]', 'g')

function slugify(text, maxWords = 5) {
  const base = (text || 'sin-titulo')
    .normalize('NFD').replace(DIACRITICS_RE, '') // quita tildes
    .toLowerCase()
    .trim()
    .split(/\s+/)
    .slice(0, maxWords)
    .join(' ')
    .replace(/[^a-z0-9\s-]/g, '')
    .replace(/\s+/g, '-')
    .replace(/-+/g, '-')
    .replace(/^-|-$/g, '')
  return base || 'sin-titulo'
}

function todayStamp() {
  const d = new Date()
  const pad = n => String(n).padStart(2, '0')
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())}`
}

// [categoria]-[tipo]-[fecha]-[titulo-corto].ext
export function buildFileName({ categoryId, typeId, titulo, extension }) {
  const parts = [categoryId, typeId, todayStamp(), slugify(titulo)]
  return `${parts.join('-')}.${extension}`
}
