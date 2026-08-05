import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import { AppProvider } from './context/AppContext'
import { AuthProvider } from './context/AuthContext'
import Layout from './components/Layout'
import Login from './pages/Login'
import Dashboard from './pages/Dashboard'
import Marketplace from './pages/Marketplace'
import Projects from './pages/Projects'
import Traceability from './pages/Traceability'
import Calculator from './pages/Calculator'
import Backoffice from './pages/Backoffice'
import Plants from './pages/Plants'
import Asistente from './pages/Asistente'
import AssetFactory from './pages/AssetFactory'

export default function App() {
  return (
    <AuthProvider>
      <AppProvider>
        <BrowserRouter>
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/" element={<Layout />}>
              <Route index element={<Navigate to="/dashboard" replace />} />
              <Route path="dashboard" element={<Dashboard />} />
              <Route path="marketplace" element={<Marketplace />} />
              <Route path="projects" element={<Projects />} />
              <Route path="traceability" element={<Traceability />} />
              <Route path="calculator" element={<Calculator />} />
              <Route path="backoffice" element={<Backoffice />} />
              <Route path="plants" element={<Plants />} />
              <Route path="asistente" element={<Asistente />} />
              <Route path="assets" element={<AssetFactory />} />
            </Route>
            <Route path="*" element={<Navigate to="/login" replace />} />
          </Routes>
        </BrowserRouter>
      </AppProvider>
    </AuthProvider>
  )
}
