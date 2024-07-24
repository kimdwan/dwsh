import "./assets/css/SignUpTwo.css"
import { SignUpTwoFooter, SignUpTwoHeader, SignUpTwoTermForm } from "./components"
import { useSignUpTwoInitValueHook } from "./hooks"

export const SignUpTwo = () => {
  const { terms, setTerms, type } = useSignUpTwoInitValueHook()

  return (
    <div className = "SignUpTwoContainer">
      
      {/* 가장 이용약간에 대한 설명이 들어있는 컴퍼넌트 */}
      <SignUpTwoHeader setTerms = { setTerms } />

      {/* 이용약간을 체크유무에 따른 컴퍼넌트 */}
      <SignUpTwoTermForm setTerms = { setTerms } />

      {/* 이용약간의 마지막 부분을 담당  */}
      <SignUpTwoFooter terms = {terms} type = {type} />

    </div>
  )
}