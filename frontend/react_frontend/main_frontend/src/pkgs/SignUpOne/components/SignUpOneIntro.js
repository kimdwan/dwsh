import  welcome  from "../assets/img/welcome.webp"

export const SignUpOneIntro = () => {
  return (
    <div className = "SignUpOneIntroContainer">

      {/* 환영을 뜻하는 이미지 */}
      <div className = "SignUpOneWelcomeImg">
        <img className = "SignUpOneWelcomeImgValue" src = {welcome} alt = "환영의 이미지" />
      </div>

      {/* 환영을 뜻하는 문구 */}
      <div className = "SignUpOneWelcomeText">
        <h1 className = "SignUpOneWelcomeTextValue">저희 홈페이지에 회원가입 해주셔서 감사합니다.</h1>
      </div>

    </div>
  )
}