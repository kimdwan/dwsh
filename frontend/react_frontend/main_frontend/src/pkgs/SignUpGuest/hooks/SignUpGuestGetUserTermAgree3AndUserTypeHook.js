import { useEffect, useState } from "react"

export const useSignUpGuestGetUserTermAgree3AndUserTypeHook = () => {
  const [ termAgree3, setTermAgree3 ] = useState(false)
  const [ userType, setUserType ] = useState("")
  useEffect(() => {
    const current_url = window.location.href
    const urlList = current_url.split("/")
    setTermAgree3(term => {
      urlList[urlList.length - 1] === "term=true" && (term = true)
      return term
    })

    setUserType(urlList[urlList.length - 3])
  }, [ ])

  return { termAgree3, userType }
}