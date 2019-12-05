import React, { FunctionComponent } from "react";
import { RouteComponentProps, navigate, Link } from "@reach/router";
import {
  PageHeader,
  Row,
  Col,
  Layout,
  Card,
  Avatar,
  Tabs,
  Divider,
  Table
} from "antd";
import routes from "../../routes";
import { Padding, PaddingTop } from "../../componets/Padding";
import { Center } from "../../componets/Center";
import { Title } from "../../componets/Heading";
import { useQuery } from "@apollo/react-hooks";
import { GET_SELLERS } from "../../gql/queries";
import { renderCardContent } from "../../componets/Card";
const { TabPane } = Tabs;

function callback(key: any) {
  console.log(key);
}

function truncate(str: string) {
  return str.length > 10 ? str.substring(0, 7) + "..." : str;
}

const columns = [
  {
    title: "ID",
    dataIndex: "id",
    render: (text: any) => (
      <Link to={`/directory/seller/${text}`}>{truncate(text)}</Link>
    )
  },
  {
    title: "Name",
    dataIndex: "name"
  },

  {
    title: "Description",
    dataIndex: "description"
  },
  {
    title: "Created At",
    dataIndex: "createdAt"
  }
];

const SellerDetails: FunctionComponent<RouteComponentProps> = (props: any) => {
  const { loading, error, data } = useQuery(GET_SELLERS, {
    variables: {
      id: props.id
    }
  });

  function getSeller() {
    if (error || loading) {
      return {
        user: {}
      };
    }
    return data.sellers.list[0];
  }

  const hasProducts = data !== undefined;
  const content = {
    title: "Total Sales",
    total: data ? data.sellers.count : 0
  };

  const content_products = {
    title: "Total Products",
    total: data ? getSeller().products.length : 0
  };

  const content_product2 = {
    title: "Total Refunds",
    total: data ? getSeller().products.length : 0
  };

  const content_product3 = {
    title: "Total Sales IN KSH",
    total: data ? getSeller().products.length : 0
  };


  function createColumnData() {
    return getSeller().products.map((item: any) => ({
      key: item.id,
      id: item.id,
      name: item.name,
      description: item.description,
      createdAt: new Date(item.createdAt).toLocaleString()
    }));
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
                  <PaddingTop />
                  <Row gutter={16}>
                    <Col className="gutter-row" span={6}>
                      {renderCardContent(content, "1")}
                    </Col>

                    <Col className="gutter-row" span={6}>
                      {renderCardContent(content_products, "2")}
                    </Col>


                    <Col className="gutter-row" span={6}>
                      {renderCardContent(content_product2, "3")}
                    </Col>


                    <Col className="gutter-row" span={6}>
                      {renderCardContent(content_product3, "4")}
                    </Col>
                  </Row>

                  <PaddingTop />
                </TabPane>
                <TabPane tab="Products" key="2">
                  <PaddingTop />
                  <Table
                    loading={loading}
                    // rowSelection={rowSelection}
                    columns={columns}
                    dataSource={hasProducts ? createColumnData() : null}
                  />
                </TabPane>
                <TabPane tab="Orders" key="3">
                 Orders
                </TabPane>
                <TabPane tab="Taxs" key="4">
                 
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
                <Title> {getSeller().name} </Title>

                <Divider />
                <Title> Contact Info </Title>
                <p> Email </p>
                {getSeller().user.email}

                <p> Phone </p>
                {getSeller().phone}

                <Divider />
                <Title> Account Info </Title>
                <p> Member Since </p>
                {new Date(getSeller().user.createdAt).toUTCString()}
                <p> Last Login </p>
                {new Date(getSeller().user.lastLogin).toUTCString()}
                <p> Status </p>
              </Center>
            </Card>
          </Padding>
        </Col>
      </Row>
    </>
  );
};

export default SellerDetails;
