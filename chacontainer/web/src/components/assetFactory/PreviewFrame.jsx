// Envuelve un nodo de tamaño fijo (los templates se renderizan siempre a su
// tamaño real en px) y lo escala visualmente con CSS transform para que
// quepa en el panel de vista previa. offsetWidth/offsetHeight del nodo no
// cambian por el transform, así que el mismo DOM sirve para exportar PNG/PDF.
export default function PreviewFrame({ width, height, maxWidth = 420, children }) {
  const scale = Math.min(1, maxWidth / width)
  return (
    <div
      style={{
        width: width * scale,
        height: height * scale,
        overflow: 'hidden',
        borderRadius: 12,
        border: '1px solid #1f2937',
        flexShrink: 0,
      }}
    >
      <div style={{ width, height, transform: `scale(${scale})`, transformOrigin: 'top left' }}>
        {children}
      </div>
    </div>
  )
}
