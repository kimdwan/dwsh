import { useState, useCallback, useEffect, useRef } from "react"

export const useSignUpTwoClickAllCheckHook = ( setTerms ) => {
  const [ allCheckValue, setAllCheckValue ] = useState(false)
  const checkBox = useRef(null)

  const clickAllTermCheckBtn = useCallback((event) => {
    if (event.target.className === "SignUpTwoHeaderAllCheckInputBox") {
      setAllCheckValue(val => !val)
      checkBox.current = document.querySelectorAll(".SignUpTwoTermFormHeaderInput")
      if ( checkBox.current ) {
        Array.from(checkBox.current).forEach((inputDiv) => {
          inputDiv.checked = event.target.checked
        })
      }
    }
  }, [ ])

  useEffect(() => {
    window.addEventListener("input", clickAllTermCheckBtn)
    return () => {
      window.removeEventListener("input", clickAllTermCheckBtn)
    }
  }, [ clickAllTermCheckBtn ]) 

  return { allCheckValue }
}