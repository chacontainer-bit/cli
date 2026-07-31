import PostCuadrado from './social/PostCuadrado'
import QuoteCard from './social/QuoteCard'
import BannerHorizontal from './social/BannerHorizontal'
import HistoriaVertical from './social/HistoriaVertical'
import PropuestaComercial from './comercial/PropuestaComercial'
import OnePagerServicio from './comercial/OnePagerServicio'
import DeckCorto from './presentaciones/DeckCorto'
import PortadaPresentacion from './presentaciones/PortadaPresentacion'
import LandingPreview from './landing/Landing'
import { buildLandingHtml } from '../../../lib/assetFactory/landingHtml'

// Registro de templates reutilizables. Al agregar un nuevo tipo de asset al
// catálogo (catalog.js), registrar aquí su template:
// - 'single': componente forwardRef con un único nodo a exportar (PNG o PDF de 1 página)
// - 'multi': componente que recibe { data, registerPage(index, node) } para PDFs multi-página
// - 'html': componente de preview (iframe) + función que arma el HTML exportable
export const TEMPLATE_REGISTRY = {
  'post-cuadrado': { kind: 'single', Component: PostCuadrado },
  'quote-card': { kind: 'single', Component: QuoteCard },
  'banner-horizontal': { kind: 'single', Component: BannerHorizontal },
  'historia-vertical': { kind: 'single', Component: HistoriaVertical },
  'propuesta-comercial': { kind: 'multi', Component: PropuestaComercial },
  'one-pager-servicio': { kind: 'single', Component: OnePagerServicio },
  'deck-corto': { kind: 'multi', Component: DeckCorto },
  'portada-presentacion': { kind: 'single', Component: PortadaPresentacion },
  'landing-campana': { kind: 'html', Component: LandingPreview, buildHtml: buildLandingHtml },
}

export function getTemplate(typeId) {
  return TEMPLATE_REGISTRY[typeId] || null
}
