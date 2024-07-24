import { useEffect, useCallback } from "react"
import { useNavigate } from "react-router-dom"

export const useSignUpOneGoTermHook = () => {
  const navigate = useNavigate()
  const clickGoSignUpTermBtn = useCallback((event) => {
    if (event.target.className === "SignUpOneWhoRouteDsValue") {
      navigate("/signup/term/ds/")
    } else if (event.target.className === "SignUpOneWhoRouteGuestValue") {
      navigate("/signup/term/guest/")
    }
  }, [ navigate ])

  useEffect(() => {
    window.addEventListener("click", clickGoSignUpTermBtn)
    return () => {
      window.removeEventListener("click", clickGoSignUpTermBtn)
    }
  }, [ clickGoSignUpTermBtn ])
}