const Error: React.FC<{ code: number, message: string }> = ({ code, message }) => {

    return (
        <div>
            <h1>Error {code}</h1>
            <h2>{message}</h2>
        </div>
    )
}

export default Error