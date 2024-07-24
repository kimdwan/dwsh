import { useCallback, useEffect } from "react"
import { useNavigate } from "react-router-dom"

export const useSignUpTwoClickNextBtnHook = (terms, type) => {
  const navigate = useNavigate()

  const clickNextBtn = useCallback((event) => {
    if ( event.target.className === "SignUpTwoFooterBtn" ) {
      if ( !terms[0] || !terms[1] ) {
        alert("필수1, 필수2는 체크해주셔야 합니다.")
      } else {
        const want_url = `/signup/term/${type}`
        if (terms[2]) {
          const url = `${want_url}/&/term=true`
          navigate(url)
        } else {
          const url = `${want_url}/&/term=false`
          navigate(url)
        }
      }
    }
  }, [ terms, type, navigate ])
  
  useEffect(() => {
    window.addEventListener("click", clickNextBtn)
    return () => {
      window.removeEventListener("click", clickNextBtn)
    }
  }, [ clickNextBtn ])

}