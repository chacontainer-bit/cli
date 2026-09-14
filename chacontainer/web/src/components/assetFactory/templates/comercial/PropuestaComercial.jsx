import BrandMark from '../BrandMark'
import ScaledPage from '../ScaledPage'
import { BRAND, PAGE_A4 } from '../../../../lib/assetFactory/brand'

const pageStyle = {
  background: BRAND.surfaceDark,
  fontFamily: BRAND.fontFamily,
  boxSizing: 'border-box',
  position: 'relative',
  overflow: 'hidden',
}

function today() {
  return new Date().toLocaleDateString('es-MX', { day: 'numeric', month: 'long', year: 'numeric' })
}

// registerPage(index, domNode) permite exportar múltiples páginas a un solo
// PDF. previewScale controla el zoom visual del preview sin afectar el
// tamaño real capturado (ver ScaledPage).
export default function PropuestaComercial({ data, registerPage, previewScale = 1 }) {
  const { titulo, subtitulo, datosClave, imagen } = data
  const items = (datosClave || []).filter(d => d.label || d.value)
  const { width, height } = PAGE_A4

  return (
    <>
      {/* Página 1 — Portada */}
      <ScaledPage width={width} height={height} scale={previewScale} registerRef={el => registerPage(0, el)} style={pageStyle}>
        {imagen && (
          <>
            <img src={imagen} alt="" style={{ position: 'absolute', inset: 0, width: '100%', height: '100%', objectFit: 'cover', opacity: 0.35 }} />
            <div style={{ position: 'absolute', inset: 0, background: `linear-gradient(180deg, ${BRAND.surfaceDark}66 0%, ${BRAND.surfaceDark} 85%)` }} />
          </>
        )}
        <div style={{ position: 'relative', height: '100%', display: 'flex', flexDirection: 'column', justifyContent: 'space-between', padding: '64px 60px' }}>
          <BrandMark size={26} textSize={17} />
          <div>
            <p style={{ color: BRAND.color, fontSize: 15, fontWeight: 700, letterSpacing: '0.12em', textTransform: 'uppercase', margin: '0 0 16px' }}>
              Propuesta comercial
            </p>
            <div style={{ width: 64, height: 6, background: BRAND.color, borderRadius: 3, marginBottom: 24 }} />
            <h1 style={{ color: BRAND.textLight, fontSize: 46, fontWeight: 800, lineHeight: 1.15, margin: 0 }}>
              {titulo || 'Título de la propuesta'}
            </h1>
            {subtitulo && <p style={{ color: BRAND.textMuted, fontSize: 17, lineHeight: 1.5, marginTop: 18, maxWidth: 520 }}>{subtitulo}</p>}
          </div>
          <p style={{ color: BRAND.textFaint, fontSize: 13, margin: 0 }}>{today()}</p>
        </div>
      </ScaledPage>

      {/* Página 2 — Alcance / Detalle */}
      <ScaledPage width={width} height={height} scale={previewScale} registerRef={el => registerPage(1, el)} style={pageStyle}>
        <div style={{ height: '100%', display: 'flex', flexDirection: 'column', padding: '56px 60px', boxSizing: 'border-box' }}>
          <BrandMark size={20} textSize={13} />
          <h2 style={{ color: BRAND.textLight, fontSize: 28, fontWeight: 800, margin: '40px 0 8px' }}>Alcance de la propuesta</h2>
          {subtitulo && <p style={{ color: BRAND.textMuted, fontSize: 15, lineHeight: 1.6, margin: '0 0 32px', maxWidth: 600 }}>{subtitulo}</p>}

          <div style={{ border: `1px solid ${BRAND.border}`, borderRadius: 12, overflow: 'hidden' }}>
            <div style={{ display: 'flex', background: BRAND.surfaceCard, padding: '14px 20px' }}>
              <span style={{ flex: 1, color: BRAND.color, fontSize: 13, fontWeight: 700, textTransform: 'uppercase', letterSpacing: '0.06em' }}>Concepto</span>
              <span style={{ width: 180, color: BRAND.color, fontSize: 13, fontWeight: 700, textTransform: 'uppercase', letterSpacing: '0.06em', textAlign: 'right' }}>Valor</span>
            </div>
            {items.length === 0 && (
              <div style={{ padding: '20px', color: BRAND.textFaint, fontSize: 14 }}>Sin ítems agregados</div>
            )}
            {items.map((it, i) => (
              <div
                key={i}
                style={{
                  display: 'flex',
                  padding: '16px 20px',
                  borderTop: `1px solid ${BRAND.border}`,
                  background: i % 2 === 0 ? BRAND.surfaceDark : BRAND.surface,
                }}
              >
                <span style={{ flex: 1, color: BRAND.textLight, fontSize: 15 }}>{it.label || '—'}</span>
                <span style={{ width: 180, color: BRAND.textLight, fontSize: 15, fontWeight: 600, textAlign: 'right' }}>{it.value || '—'}</span>
              </div>
            ))}
          </div>
        </div>
      </ScaledPage>

      {/* Página 3 — Cierre */}
      <ScaledPage width={width} height={height} scale={previewScale} registerRef={el => registerPage(2, el)} style={pageStyle}>
        <div style={{ height: '100%', display: 'flex', flexDirection: 'column', justifyContent: 'space-between', padding: '64px 60px', boxSizing: 'border-box' }}>
          <BrandMark size={20} textSize={13} />
          <div style={{ textAlign: 'center' }}>
            <div style={{ width: 64, height: 6, background: BRAND.color, borderRadius: 3, margin: '0 auto 28px' }} />
            <h2 style={{ color: BRAND.textLight, fontSize: 36, fontWeight: 800, margin: '0 0 12px' }}>¡Conversemos!</h2>
            <p style={{ color: BRAND.textMuted, fontSize: 16, maxWidth: 440, margin: '0 auto' }}>
              Quedamos atentos a tus comentarios para avanzar con esta propuesta.
            </p>
          </div>
          <p style={{ color: BRAND.textFaint, fontSize: 13, textAlign: 'center', margin: 0 }}>{BRAND.name} · {BRAND.tagline}</p>
        </div>
      </ScaledPage>
    </>
  )
}
