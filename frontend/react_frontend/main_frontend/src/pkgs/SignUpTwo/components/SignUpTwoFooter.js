import { useSignUpTwoClickNextBtnHook } from "../hooks"

export const SignUpTwoFooter = ({ terms, type }) => {
  useSignUpTwoClickNextBtnHook(terms, type)

  return (
    <div className = "SignUpTwoFooterContainer">

      <button className = "SignUpTwoFooterBtn">
        다음
      </button>

    </div>
  )
}