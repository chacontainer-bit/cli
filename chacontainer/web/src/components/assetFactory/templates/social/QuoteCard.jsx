import { forwardRef } from 'react'
import { Quote } from 'lucide-react'
import BrandMark from '../BrandMark'
import { BRAND } from '../../../../lib/assetFactory/brand'

const QuoteCard = forwardRef(function QuoteCard({ data }, ref) {
  const { titulo, autor } = data
  return (
    <div
      ref={ref}
      style={{
        width: 1080,
        height: 1080,
        background: BRAND.surface,
        fontFamily: BRAND.fontFamily,
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between',
        padding: '80px 96px',
        boxSizing: 'border-box',
      }}
    >
      <Quote size={64} color={BRAND.color} fill={BRAND.color} strokeWidth={0} />

      <p
        style={{
          color: BRAND.textLight,
          fontSize: 56,
          fontWeight: 700,
          lineHeight: 1.28,
          margin: 0,
        }}
      >
        {titulo || 'Escribe aquí la frase o cita destacada.'}
      </p>

      <div>
        {autor && (
          <p style={{ color: BRAND.color, fontSize: 26, fontWeight: 600, margin: '0 0 32px' }}>{autor}</p>
        )}
        <BrandMark size={26} textSize={17} />
      </div>
    </div>
  )
})

export default QuoteCard
