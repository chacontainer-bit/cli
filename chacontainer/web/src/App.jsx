import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import { AppProvider } from './context/AppContext'
import Layout from './components/Layout'
import MarketingLanding from './pages/MarketingLanding'
import Login from './pages/Login'
import Dashboard from './pages/Dashboard'
import Marketplace from './pages/Marketplace'
import Projects from './pages/Projects'
import Traceability from './pages/Traceability'
import Calculator from './pages/Calculator'
import Backoffice from './pages/Backoffice'
import Asistente from './pages/Asistente'
import AssetFactory from './pages/AssetFactory'

export default function App() {
  return (
    <AppProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<MarketingLanding />} />
          <Route path="/login" element={<Login />} />
          <Route path="/app" element={<Layout />}>
            <Route index element={<Navigate to="/app/dashboard" replace />} />
            <Route path="dashboard" element={<Dashboard />} />
            <Route path="marketplace" element={<Marketplace />} />
            <Route path="projects" element={<Projects />} />
            <Route path="traceability" element={<Traceability />} />
            <Route path="calculator" element={<Calculator />} />
            <Route path="backoffice" element={<Backoffice />} />
            <Route path="asistente" element={<Asistente />} />
            <Route path="assets" element={<AssetFactory />} />
          </Route>
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </BrowserRouter>
    </AppProvider>
  )
}
