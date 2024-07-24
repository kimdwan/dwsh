import { useCallback, useState } from "react"

export const useSignUpOneOnMouseHook = () => {
  const [ choiceDs, setChoiceDs ] = useState(false)
  const [ choiceGuest, setChoiceGuest ] = useState(false)
  
  const choiceUserType = useCallback((event) => {
    if (event.target.className === "SignUpOneWhoRouteDsValue") {
      setChoiceDs(true)
    } else if (event.target.className === "SignUpOneWhoRouteGuestValue") {
      setChoiceGuest(true)
    }
  }, [ ])
  
  const leaveUserType = useCallback((event) => {
    if (event.target.className === "SignUpOneWhoRouteDsValue") {
      setChoiceDs(false)
    } else if (event.target.className === "SignUpOneWhoRouteGuestValue") {
      setChoiceGuest(false)
    }
  }, [ ])


  return { choiceDs, choiceGuest, choiceUserType, leaveUserType }
}