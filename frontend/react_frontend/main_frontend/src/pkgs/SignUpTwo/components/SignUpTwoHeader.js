import { useSignUpTwoClickAllCheckHook, useSignUpTwoMakeTermDataHook } from "../hooks"

export const SignUpTwoHeader = ({setTerms}) => {
  const { allCheckValue } = useSignUpTwoClickAllCheckHook()
  useSignUpTwoMakeTermDataHook(allCheckValue, setTerms)

  return (
    <div className = "SignUpTwoHeaderContainer">
      
      {/* 이용약간임을 알려주는 컴퍼넌트 */}
      <div className = "SignUpTwoHeaderIntro">
        <h1 className = "SignUpTwoHeaderIntroText">
          이용약간
        </h1>
      </div>

      {/* 전체 선택이 가능하게 함 */}
      <div className = "SignUpTwoHeaderAllCheck">
        
        <input 
          type = "checkbox"
          className = "SignUpTwoHeaderAllCheckInputBox"
          value = {allCheckValue}
        />

        <div className = "SignUpTwoHeaderAllCheckText">
          전체동의
        </div>

      </div>


    </div>
  )
}