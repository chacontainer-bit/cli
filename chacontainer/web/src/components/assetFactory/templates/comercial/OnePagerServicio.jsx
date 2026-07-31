import { forwardRef } from 'react'
import { CheckCircle2 } from 'lucide-react'
import BrandMark from '../BrandMark'
import { BRAND, PAGE_A4 } from '../../../../lib/assetFactory/brand'

const OnePagerServicio = forwardRef(function OnePagerServicio({ data }, ref) {
  const { titulo, subtitulo, datosClave, imagen } = data
  const features = (datosClave || []).filter(d => d.label || d.value)

  return (
    <div
      ref={ref}
      style={{
        width: PAGE_A4.width,
        height: PAGE_A4.height,
        background: BRAND.surfaceDark,
        fontFamily: BRAND.fontFamily,
        boxSizing: 'border-box',
        display: 'flex',
        flexDirection: 'column',
      }}
    >
      {imagen && (
        <div style={{ height: 220, position: 'relative', flexShrink: 0 }}>
          <img src={imagen} alt="" style={{ width: '100%', height: '100%', objectFit: 'cover' }} />
          <div style={{ position: 'absolute', inset: 0, background: `linear-gradient(180deg, transparent 40%, ${BRAND.surfaceDark} 100%)` }} />
        </div>
      )}

      <div style={{ padding: '40px 56px 48px', display: 'flex', flexDirection: 'column', flex: 1 }}>
        <BrandMark size={20} textSize={13} />

        <div style={{ marginTop: imagen ? 8 : 40 }}>
          <div style={{ width: 56, height: 6, background: BRAND.color, borderRadius: 3, marginBottom: 20 }} />
          <h1 style={{ color: BRAND.textLight, fontSize: 34, fontWeight: 800, margin: 0 }}>{titulo || 'Nombre del servicio'}</h1>
          {subtitulo && <p style={{ color: BRAND.textMuted, fontSize: 16, lineHeight: 1.55, marginTop: 14, maxWidth: 620 }}>{subtitulo}</p>}
        </div>

        <div style={{ marginTop: 32, display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 16, flex: 1 }}>
          {features.map((f, i) => (
            <div key={i} style={{ display: 'flex', gap: 12, alignItems: 'flex-start' }}>
              <CheckCircle2 size={20} color={BRAND.color} style={{ flexShrink: 0, marginTop: 2 }} />
              <div>
                {f.label && <p style={{ color: BRAND.textLight, fontSize: 15, fontWeight: 700, margin: 0 }}>{f.label}</p>}
                {f.value && <p style={{ color: BRAND.textMuted, fontSize: 13.5, lineHeight: 1.5, margin: '4px 0 0' }}>{f.value}</p>}
              </div>
            </div>
          ))}
          {features.length === 0 && (
            <p style={{ color: BRAND.textFaint, fontSize: 14 }}>Sin características agregadas</p>
          )}
        </div>

        <div style={{ borderTop: `1px solid ${BRAND.border}`, paddingTop: 18, marginTop: 'auto' }}>
          <p style={{ color: BRAND.textFaint, fontSize: 12.5, margin: 0 }}>{BRAND.name} · {BRAND.tagline}</p>
        </div>
      </div>
    </div>
  )
})

export default OnePagerServicio
