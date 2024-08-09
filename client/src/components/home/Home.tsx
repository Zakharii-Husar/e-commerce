import styled from 'styled-components'
import ProductsList from '../products/ProductsList'

const Home: React.FC = () => {
  return (
    <Container>
      <ProductsList />
    </Container>
  )
}

export default Home

const Container = styled.div`
max-width:1500px;
margin:0 auto;
`
