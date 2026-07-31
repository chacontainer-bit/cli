import { BRAND } from '../../../../lib/assetFactory/brand'
import { buildLandingHtml } from '../../../../lib/assetFactory/landingHtml'

export default function LandingPreview({ data }) {
  const html = buildLandingHtml(data)
  return (
    <iframe
      title="Vista previa landing page"
      srcDoc={html}
      style={{ width: 1280, height: 1600, border: 'none', background: BRAND.surfaceDark }}
    />
  )
}
