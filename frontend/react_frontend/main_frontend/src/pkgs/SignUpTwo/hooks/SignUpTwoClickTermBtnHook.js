import { useCallback, useEffect } from "react"

export const useSignUpTwoClickTermBtnHook = (setTerms) => {
  const clickTermBtn = useCallback((event) => {

    const mainInputName = "SignUpTwoTermFormHeaderInput"
    if ( event.target.className === "SignUpTwoTermFormHeaderInput" ) {
      const checkClickBoxId = event.target.id 

      Array.from(["one", "two", "three"]).forEach((name, idx) => {
        if ( checkClickBoxId === `${mainInputName}${name}` ) {
          setTerms((prev) => {
            const newTerms = [ ...Array.from(prev) ]
            newTerms[idx] = !newTerms[idx]
            return newTerms
          })           
        }
      })

    }
    
  },[ setTerms ])

  useEffect(() => {
    window.addEventListener("input", clickTermBtn)
    return () => {
      window.removeEventListener("input", clickTermBtn)
    }
  }, [ clickTermBtn ])
}