import { useEffect, useState } from "react"

export const useSignUpDsUserTerm3AndUserTypeHook = () => {
  const [ userAgree3, setUserAgree3 ] = useState(false)
  const [ userType, setUserType ] = useState("")
  useEffect(() => {
    const current_url = window.location.href
    const urlList = current_url.split("/")
    const userPick3 = urlList[urlList.length - 1]
    setUserAgree3( (term) => {
      userPick3 === "term=true" && (term = true)
      return term
    })
    setUserType(urlList[urlList.length - 3].toUpperCase())

  }, [ ])

  return { userAgree3, userType }
}