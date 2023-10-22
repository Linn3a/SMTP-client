import { useState } from 'react'
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Layout from "./components/layout"
import Login from "./pages/login"
import Contact from './pages/contact';
import Send from './pages/send'
import Box from './pages/box'
import Craft from './pages/craft'
import Edit from './pages/edit'

const queryClient = new QueryClient();
function App() {
  const [isLogin,setIslogin] = useState(false);
  
  
  return ( 
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Routes>
          {/* <Route path="/login" element={<Login />} /> */}
          <Route
            path="/"
            element={
              isLogin?
              <Layout />:
              
                <Login setIslogin={setIslogin} />
              
            }
          >
            {/* <Route path="/" element={<Statistics/>}/> */}
            <Route path="/contact" element={<Contact/>} />
            <Route path="/send" element={<Send/>} />
            <Route path="/box" element={<Box />} />
            <Route path="craft" element={<Craft/>} />
              <Route path='/craft/:id' element={<Edit />}/>
            </Route>
        </Routes>
      </BrowserRouter>
      </QueryClientProvider>
  )
}

export default App
