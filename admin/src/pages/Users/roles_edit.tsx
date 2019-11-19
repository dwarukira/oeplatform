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

const RoleEdit: FunctionComponent<RouteComponentProps> = (props: any) => {
  const { loading, data, error } = useQuery(GET_ROLE, {
    variables: {
      id: props.id
    }
  });

  function callback(key: any) {
    console.log(key);
  }

  function getRole() {
    console.log(data, "roles");
    
    if (error) {
      return {};
    }
    return data.roles.list[0];
  }

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
            {" "}
            <Title> Editing {getRole().name}' Role </Title>{" "}
          </Padding>
          <Padding>
            <Row gutter={8}>
              <Col span={8}>
                <Form layout="vertical">
                  <Collapse defaultActiveKey={["1"]} onChange={callback} bordered={false} >
                    <Panel header="Details" key="1">
                      <Form.Item label="Title">
                        <Input placeholder="title" value={getRole().name} />
                      </Form.Item>
                    </Panel>

                    <Panel header="Permissions" key="2">
  
                 

                    <Form.Item label="">
                    <Checkbox checked={true}>
                      {`${getRole().permissions[0].tag} - ${getRole().permissions[0].description}` }
                    </Checkbox>
                  </Form.Item>

                  </Panel>
                  </Collapse>

                  <Padding>
                  <Form.Item>
                    <Button type="primary">Save Changes</Button>
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

export default RoleEdit;
