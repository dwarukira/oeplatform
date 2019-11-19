import React, { FunctionComponent } from "react";
import { RouteComponentProps, navigate } from "@reach/router";
import { PageHeader, Row, Col, Layout, Card, Avatar, Tabs } from "antd";
import routes from "../../routes";
import { Padding, PaddingTop } from "../../componets/Padding";
import { Center } from "../../componets/Center";
import { Title } from "../../componets/Heading";
import { useQuery } from "@apollo/react-hooks";
import { GET_SELLERS } from "../../gql/queries";
const { TabPane } = Tabs;

function callback(key: any) {
  console.log(key);
}

const SellerDetails: FunctionComponent<RouteComponentProps> = (props: any) => {

  const { loading, error, data } = useQuery(GET_SELLERS, {
    variables: {
      id: props.id
    }
  });


  function getSeller() {
    if (error || loading) {
      return {}
    }
    return data.sellers.list[0];
  }


  return (
    <>
      <PageHeader
        onBack={() => navigate(routes.directory_sellers)}
        title=""
        subTitle="Back to sellers  list"
      />

      <Row>
        <Col span={18} push={6}>
          <Padding>
            <Card
             style={{
              minHeight: 500
            }}
            >
              <Tabs defaultActiveKey="1" onChange={callback}>
                <TabPane tab="Details" key="1">
                  Content of Tab Pane 1
                </TabPane>
                <TabPane tab="Products" key="2">
                  Content of Tab Pane 2
                </TabPane>
                <TabPane tab="Orders" key="3">
                  Content of Tab Pane 3
                </TabPane>
                <TabPane tab="Taxs" key="4">
                  Content of Tab Pane 4
                </TabPane>
              </Tabs>
              ,
            </Card>
          </Padding>
        </Col>
        <Col span={6} pull={18}>
          <Padding>
            <Card
              style={{
                minHeight: 500
              }}
            >
              <Center>
                {" "}
                <Avatar src="https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png" />
              </Center>
              <PaddingTop />
              <Center>
                <Title> { getSeller().name } </Title>
              </Center>
            </Card>
          </Padding>
        </Col>
      </Row>
    </>
  );
};

export default SellerDetails;
