

export const SignUpGuestSignUpFormFetch = async ( url, datas, setError ) => {
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
        alert("클라이언트에서 보낸 폼에 문제가 있음")
        throw new Error("클라이언트에서 보낸 폼이 잘못되었습니다.")
      } else if (response.status === 500) {
        alert("서버에 오류가 발생했습니다.")
        throw new Error("서버에 오류가 발생했습니다.")
      } else if (response.status === 510) {
        setError("email", {
          type : "manual",
          message : "이미 존재하는 이메일 입니다.",
        })
        throw new Error("이미 존재하는 이메일")
      } else {
        throw new Error(`오류가 발생했습니다 오류번호: ${response.status}`)
      }
    }

    const data = await response.json()

    return data

  } catch (err) {
    throw err
  }
}