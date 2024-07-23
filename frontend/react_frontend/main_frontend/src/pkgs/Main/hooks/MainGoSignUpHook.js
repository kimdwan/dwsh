import { useCallback, useEffect } from "react"
import { useNavigate } from "react-router-dom"

export const useMainGoSignUpHook = () => {
  const navigate = useNavigate()

  const clickGoSignUpBtn = useCallback((event) => {
    if (event.target.className === "mainGoSignUpBtn") {
      const url = "/signup/"
      navigate(url)
    }
  }, [ navigate ])

  useEffect(() => {
    window.addEventListener("click", clickGoSignUpBtn)
    return () => {
      window.removeEventListener("click", clickGoSignUpBtn)
    }
  }, [ clickGoSignUpBtn ])

}