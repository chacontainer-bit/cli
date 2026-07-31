// Catálogo de la Fábrica de Assets: categorías -> opciones de assets.
// Cada asset type declara cómo se exporta y qué campos de contenido soporta,
// para que el formulario y el motor de export sean genéricos y data-driven.
import { PAGE_A4, SLIDE_16_9 } from './brand'

export const EXPORT_KIND = {
  PNG: 'png',
  PDF: 'pdf',
  HTML: 'html',
}

export const CATEGORIES = [
  {
    id: 'redes-sociales',
    label: 'Redes Sociales',
    description: 'Piezas gráficas listas para publicar',
    types: [
      {
        id: 'post-cuadrado',
        label: 'Post cuadrado',
        description: '1080 × 1080 · Feed Instagram/LinkedIn',
        width: 1080,
        height: 1080,
        exportKind: EXPORT_KIND.PNG,
        fields: { subtitulo: true, datosClave: false, imagen: true, cta: false },
      },
      {
        id: 'quote-card',
        label: 'Quote card',
        description: '1080 × 1080 · Cita o frase destacada',
        width: 1080,
        height: 1080,
        exportKind: EXPORT_KIND.PNG,
        fields: { subtitulo: false, datosClave: false, imagen: false, cta: false, autor: true },
      },
      {
        id: 'banner-horizontal',
        label: 'Banner horizontal',
        description: '1200 × 628 · LinkedIn, blog, email',
        width: 1200,
        height: 628,
        exportKind: EXPORT_KIND.PNG,
        fields: { subtitulo: true, datosClave: false, imagen: true, cta: true },
      },
      {
        id: 'historia-vertical',
        label: 'Historia vertical',
        description: '1080 × 1920 · Instagram/Facebook Stories',
        width: 1080,
        height: 1920,
        exportKind: EXPORT_KIND.PNG,
        fields: { subtitulo: true, datosClave: true, imagen: true, cta: true },
      },
    ],
  },
  {
    id: 'comercial',
    label: 'Comercial',
    description: 'Documentos de venta y propuestas',
    types: [
      {
        id: 'propuesta-comercial',
        label: 'Oferta / Propuesta comercial',
        description: 'PDF de 3 páginas: portada, alcance y cotización',
        exportKind: EXPORT_KIND.PDF,
        pageSize: [PAGE_A4.width, PAGE_A4.height],
        fields: { subtitulo: true, datosClave: true, imagen: true, cta: false, datosClaveLabel: 'Ítems de la propuesta (concepto / valor)' },
      },
      {
        id: 'one-pager-servicio',
        label: 'One-pager de servicio',
        description: 'PDF de 1 página con features y datos de contacto',
        exportKind: EXPORT_KIND.PDF,
        pageSize: [PAGE_A4.width, PAGE_A4.height],
        fields: { subtitulo: true, datosClave: true, imagen: true, cta: false, datosClaveLabel: 'Características (título / detalle)' },
      },
    ],
  },
  {
    id: 'presentaciones',
    label: 'Presentaciones',
    description: 'Slides para reuniones y pitches',
    types: [
      {
        id: 'deck-corto',
        label: 'Deck corto (5 a 8 slides)',
        description: 'PDF 16:9 · portada + contenido + cierre',
        exportKind: EXPORT_KIND.PDF,
        pageSize: [SLIDE_16_9.width, SLIDE_16_9.height],
        fields: { subtitulo: true, datosClave: true, imagen: true, cta: false, datosClaveLabel: 'Slides de contenido (título / texto), 3 a 6 filas' },
      },
      {
        id: 'portada-presentacion',
        label: 'Portada de presentación',
        description: 'PDF 16:9 · slide único de apertura',
        exportKind: EXPORT_KIND.PDF,
        pageSize: [SLIDE_16_9.width, SLIDE_16_9.height],
        fields: { subtitulo: true, datosClave: false, imagen: true, cta: false },
      },
    ],
  },
  {
    id: 'landing',
    label: 'Landing Page',
    description: 'Página de campaña autocontenida',
    types: [
      {
        id: 'landing-campana',
        label: 'Landing page de campaña',
        description: 'Archivo .html autocontenido, listo para subir a hosting',
        exportKind: EXPORT_KIND.HTML,
        fields: { subtitulo: true, datosClave: true, imagen: true, cta: true, datosClaveLabel: 'Beneficios / features (título / detalle)' },
      },
    ],
  },
]

export function findAssetType(categoryId, typeId) {
  const category = CATEGORIES.find(c => c.id === categoryId)
  if (!category) return null
  const type = category.types.find(t => t.id === typeId)
  if (!type) return null
  return { category, type }
}
