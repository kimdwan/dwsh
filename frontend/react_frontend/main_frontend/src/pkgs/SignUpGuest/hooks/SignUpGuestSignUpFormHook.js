import * as yup from "yup"
import { yupResolver } from "@hookform/resolvers/yup"
import { useForm } from "react-hook-form"
import { SignUpGuestSignUpFormFetch } from "../functions"
import { useNavigate } from "react-router-dom"

export const useSignUpGuestSignUpFormHook = (termAgree3, userType) => {
  const navigate = useNavigate()

  const schema = yup.object({
    email : yup.string().email("이메일 형식을 지켜주시길 바랍니다.").required("이메일은 필수적으로 입력해주셔야 합니다."),
    password : yup.string().min(4, "비밀번호는 최소 4글자 입니다.").max(16, "비밀번호는 최대 16글자 입니다.").required("비밀번호는 필수적으로 입력해주셔야 합니다."),
    confirm_password : yup.string().oneOf([yup.ref("password")], "비밀번호를 확인해주세요."),
    user_name : yup.string().max(10, "이름은 최대 10글자 입니다.").required("이름은 필수적으로 입력해주셔야 합니다."),
  })

  const { register, handleSubmit, formState:{errors}, setError } = useForm({
    resolver : yupResolver(schema)
  })

  const go_backend_url = process.env.REACT_APP_GO_BACKEND_URL

  const onSubmit = async ( data ) => {

    // 비밀번호 확인 term 삭제
    delete data["confirm_password"]

    // 동의서 추가 
    Array.from(["term_agree_1", "term_agree_2", "term_agree_3"]).forEach((term_name, idx) => {
      if (idx !== 2) {
        data[term_name] = true
      } else {
        data[term_name] = termAgree3
      }
    })

    // 유저 타입 추가
    data["user_type"] = `${userType}`.toUpperCase()

    const url = `${go_backend_url}/user/signup`
    
    // 본격적으로 확인하는 루트
    const response =  await SignUpGuestSignUpFormFetch(url, data, setError)
    if (response) {
      alert(response["message"])
      navigate("/")
    }

    
  }

  return { register, handleSubmit, errors, onSubmit }
}