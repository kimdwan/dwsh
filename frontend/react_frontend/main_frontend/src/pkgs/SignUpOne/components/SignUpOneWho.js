import { useSignUpOneGoTermHook, useSignUpOneOnMouseHook } from "../hooks"


export const SignUpOneWho = () => {
  const { choiceDs, choiceGuest, choiceUserType, leaveUserType } = useSignUpOneOnMouseHook()
  useSignUpOneGoTermHook()

  return (
    <div className = "SignUpOneWhoContainer">

      {/* 회원가입 종류를 물어보는 박스 */}
      <div className = "SignUpOneWhoIntro">
        <h1 className = "SignUpOneWhoIntroValue">어느 회원으로 회원가입 하실건가요?</h1>
      </div>

      {/* 회원가입의 경로를 결정하는 박스 */}
      <div className = "SignUpOneWhoRoute">

        {/* DS경로 */}
        <div className = "SignUpOneWhoRouteDs">
          
          <div className = "SignUpOneWhoRouteDsText">
            <h1 className = "SignUpOneWhoRouteDsValue" onMouseEnter={choiceUserType} onMouseLeave={leaveUserType}>DS</h1>
          </div>

          <div className = "SignUpOneWhoRouteDsHidden">
            <p className = "SignUpOneWhoRouteDsHiddenValue">{choiceDs && "자기야?"}</p>
          </div>
        </div>

        {/* GUEST 경로 */}
        <div className = "SignUpOneWhoRouteGuest">

          <div className = "SignUpOneWhoRouteGuestText">
            <h1 className = "SignUpOneWhoRouteGuestValue" onMouseEnter={choiceUserType} onMouseLeave={leaveUserType}>Guest</h1>
          </div>

          <div className = "SignUpOneWhoRouteGuestHidden">
            <p className = "SignUpOneWhoRouteGuestHiddenValue">{choiceGuest && "지인분" }</p>
          </div>  
        </div>

      </div>

    </div>
  )
}