import BrandMark from '../BrandMark'
import ScaledPage from '../ScaledPage'
import { BRAND, SLIDE_16_9 } from '../../../../lib/assetFactory/brand'

const slideBase = {
  background: BRAND.surfaceDark,
  fontFamily: BRAND.fontFamily,
  boxSizing: 'border-box',
  position: 'relative',
  overflow: 'hidden',
  padding: '56px 80px',
  display: 'flex',
  flexDirection: 'column',
}

function SlideFooter({ index, total }) {
  return (
    <div style={{ position: 'absolute', bottom: 32, right: 80, color: BRAND.textFaint, fontSize: 13, fontFamily: BRAND.fontFamily }}>
      {String(index + 1).padStart(2, '0')} / {String(total).padStart(2, '0')}
    </div>
  )
}

// Entre 3 y 6 slides de contenido a partir de "datos clave", para mantener
// el deck entre 5 y 8 slides totales (portada + contenido + cierre).
function buildContentSlides(datosClave) {
  const provided = (datosClave || []).filter(d => d.label || d.value)
  const slides = provided.slice(0, 6)
  while (slides.length < 3) {
    slides.push({ label: `Punto ${slides.length + 1}`, value: 'Agrega el contenido de este slide.' })
  }
  return slides
}

// registerPage(index, domNode) permite exportar todos los slides a un solo
// PDF. previewScale controla el zoom visual del preview sin afectar el
// tamaño real capturado (ver ScaledPage).
export default function DeckCorto({ data, registerPage, previewScale = 1 }) {
  const { titulo, subtitulo, datosClave, imagen } = data
  const contentSlides = buildContentSlides(datosClave)
  const total = contentSlides.length + 2
  const { width, height } = SLIDE_16_9

  return (
    <>
      {/* Slide de portada */}
      <ScaledPage
        width={width}
        height={height}
        scale={previewScale}
        registerRef={el => registerPage(0, el)}
        style={{ ...slideBase, justifyContent: 'space-between' }}
      >
        {imagen && (
          <>
            <img src={imagen} alt="" style={{ position: 'absolute', inset: 0, width: '100%', height: '100%', objectFit: 'cover', opacity: 0.28 }} />
            <div style={{ position: 'absolute', inset: 0, background: `linear-gradient(180deg, ${BRAND.surfaceDark}55 0%, ${BRAND.surfaceDark}f2 100%)` }} />
          </>
        )}
        <div style={{ position: 'relative' }}>
          <BrandMark size={26} textSize={17} />
        </div>
        <div style={{ position: 'relative' }}>
          <div style={{ width: 72, height: 7, background: BRAND.color, borderRadius: 4, marginBottom: 24 }} />
          <h1 style={{ color: BRAND.textLight, fontSize: 54, fontWeight: 800, lineHeight: 1.1, margin: 0, maxWidth: 900 }}>
            {titulo || 'Título de la presentación'}
          </h1>
          {subtitulo && <p style={{ color: BRAND.textMuted, fontSize: 22, marginTop: 18, maxWidth: 760 }}>{subtitulo}</p>}
        </div>
        <SlideFooter index={0} total={total} />
      </ScaledPage>

      {/* Slides de contenido */}
      {contentSlides.map((item, i) => (
        <ScaledPage
          key={i}
          width={width}
          height={height}
          scale={previewScale}
          registerRef={el => registerPage(i + 1, el)}
          style={{ ...slideBase, justifyContent: 'center' }}
        >
          <div style={{ position: 'absolute', top: 56, left: 80 }}>
            <BrandMark size={18} textSize={12} />
          </div>
          <div style={{ width: 56, height: 6, background: BRAND.color, borderRadius: 3, marginBottom: 24 }} />
          <h2 style={{ color: BRAND.textLight, fontSize: 40, fontWeight: 800, margin: 0, maxWidth: 920 }}>{item.label}</h2>
          <p style={{ color: BRAND.textMuted, fontSize: 22, lineHeight: 1.5, marginTop: 20, maxWidth: 880 }}>{item.value}</p>
          <SlideFooter index={i + 1} total={total} />
        </ScaledPage>
      ))}

      {/* Slide de cierre */}
      <ScaledPage
        width={width}
        height={height}
        scale={previewScale}
        registerRef={el => registerPage(contentSlides.length + 1, el)}
        style={{ ...slideBase, justifyContent: 'center', alignItems: 'center', textAlign: 'center' }}
      >
        <div style={{ width: 56, height: 6, background: BRAND.color, borderRadius: 3, marginBottom: 24 }} />
        <h2 style={{ color: BRAND.textLight, fontSize: 40, fontWeight: 800, margin: 0 }}>Gracias</h2>
        <p style={{ color: BRAND.textMuted, fontSize: 18, marginTop: 12 }}>{BRAND.name} · {BRAND.tagline}</p>
        <SlideFooter index={contentSlides.length + 1} total={total} />
      </ScaledPage>
    </>
  )
}
