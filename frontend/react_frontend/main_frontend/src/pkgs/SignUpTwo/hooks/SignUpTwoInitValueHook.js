import { useEffect, useState } from "react"

export const useSignUpTwoInitValueHook = () => {
  const [ terms, setTerms ] = useState([false,false,false])
  const [ type, setType ] = useState("")
  useEffect(() => {
    const current_url = window.location.href
    const typeTextLists = current_url.split("/")
    const typeText = typeTextLists[ typeTextLists.length - 2 ]
    setType(typeText)

  }, [ ])

  return { terms, setTerms, type }
}