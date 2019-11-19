import React, { FunctionComponent } from "react";
import { RouteComponentProps, navigate, Link } from "@reach/router";
import {
  PageHeader,
  Spin,
  Form,
  Input,
  Button,
  Checkbox,
  Row,
  Col,
  Collapse
} from "antd";
import { Title } from "../../componets/Heading";
import { useQuery } from "@apollo/react-hooks";
import { GET_USER, GET_ROLE } from "../../gql/queries";
import { Padding } from "../../componets/Padding";

const { Panel } = Collapse;


const RoleCreate: FunctionComponent<RouteComponentProps> = (props: any) => {
  const { loading, data, error } = useQuery(GET_ROLE);

  // function callback(key: any) {
  //   console.log(key);
  // }

  // function getRole() {
  //   console.log(data, "roles");

  //   if (error) {
  //     return {};
  //   }
  //   return data.roles.list[0];
  // }

  return (
    <>
      <PageHeader
        onBack={() => navigate(`/users/roles`)}
        title=""
        subTitle="Back to user roles list"
      />
      {loading ? (
        <>
          <Spin />
        </>
      ) : (
        <>
          <Padding>
            <Title> Add a New Role </Title>
          </Padding>
          <Padding>
            <Row gutter={8}>
              <Col span={8}>
                <Form layout="vertical">
                  <Collapse defaultActiveKey={["1", "2"]} bordered={false}>
                    <Panel header="Details" key="1">
                      <Form.Item label="Title">
                        <Input placeholder="title" />
                      </Form.Item>
                    </Panel>

                    <Panel header="Permissions" key="2">
                      <Form.Item label="">
                        <Checkbox checked={false}>
                          {`Seller.Create - Can create a new seller`}
                        </Checkbox>
                      </Form.Item>

                      <Form.Item label="">
                        <Checkbox checked={false}>
                          {`Seller.Create - Can create a new seller`}
                        </Checkbox>
                      </Form.Item>



                    </Panel>
                  </Collapse>

                  <Padding>
                    <Form.Item>
                      <Button type="primary">Add new role</Button>
                    </Form.Item>
                  </Padding>
                </Form>
              </Col>
            </Row>
          </Padding>
        </>
      )}
    </>
  );
};

export default RoleCreate;
