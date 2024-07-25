

export const SignUpDsSignUpFormFetch = async ( url, datas, setError ) => {
  try {
    const response = await fetch(url, {
      method : "POST",
      headers : {
        "Content-Type" : "application/json",
        "X-Requested-With" : "XMLHttpRequest",
      },
      body : JSON.stringify(datas),
      credentials : "include",
    })

    if (!response.ok) {
      if (response.status === 400) {
        alert("클라이언트에 보낸 폼이 잘못된 형식 입니다")
        throw new Error("클라이언트에서 보낸 폼을 파싱하는데 오류가 발생했습니다")
      } else if (response.status === 502) {
        alert("필수 사항 두개는 동의를 해주셔야 합니다.")
        throw new Error("필수 사항 두개는 동의를 하셔야 합니다.")
      } else if (response.status === 406) {
        setError("our_first_day", {
          type : "manual",
          message : "첫번째 날을 맞추지 못함",
        })
        throw new Error("첫번째로 만난 날을 맞추지 못함")
      } else if (response.status === 401) {
        setError("secret_key", {
          type : "manual",
          message : "비밀키가 잘못되었습니다",
        })
        throw new Error("비밀키가 잘못됨")
      } else if (response.status === 510) {
        setError("email", {
          type : "manual",
          message : "이미 존재하는 이메일 입니다.",
        })
        throw new Error("이미 존재하는 이메일 입니다.")
      } else if (response.status === 500) {
        alert("서버에 오류가 발생했습니다.")  
        throw new Error("서버에 오류가 발생했습니다.")
      } else {
        alert("알수 없는 오류 발생")
        throw new Error(`오류가 발생했습니다 오류번호: ${response.status}`)
      }
    }

    const backend_data = await response.json()

    return backend_data

  } catch (err) {
    throw err
  }
}