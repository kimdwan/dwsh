import { useCallback, useEffect } from "react"

export const useSignUpTwoMakeTermDataHook = (allCheckValue, setTerms) => {
  const makeTermData = useCallback((val) => {
    const termLists = []
    for ( let i = 0; i < 3; i++ ) {
      termLists.push(val)
    }
    return termLists
  },[ ])

  useEffect(() => {
    allCheckValue ? setTerms(makeTermData(true)) : setTerms(makeTermData(false))
  }, [ allCheckValue, makeTermData, setTerms ])
}