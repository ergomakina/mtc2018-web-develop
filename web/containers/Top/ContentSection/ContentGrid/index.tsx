import * as React from 'react';
import { MiniGrid } from '../../../../components';
import ContentGridItem, {
  CONTENT_GRID_SESSION_FRAGMENT
} from './ContentGridItem';

import gql from 'graphql-tag';
import { ContentGridFragment } from '../../../../graphql/generated/ContentGridFragment';

export const CONTENT_GRID_FRAGMENT = gql`
  fragment ContentGridFragment on Query {
    sessionList {
      nodes {
        ...ContentGridSessionFragment
      }
    }
  }

  ${CONTENT_GRID_SESSION_FRAGMENT}
`;

interface Props {
  data: ContentGridFragment;
  onClickItem: (sessionId: number) => void;
}

class ContentGrid extends React.Component<Props> {
  public render() {
    const { data, onClickItem } = this.props;
    const { sessionList } = data;
    return (
      <MiniGrid minColumnWidth={360}>
        {sessionList.nodes!.map((session, index) => (
          <ContentGridItem
            key={`${session.id}_${index}`}
            session={session}
            index={index}
            onClick={onClickItem}
          />
        ))}
      </MiniGrid>
    );
  }
}

export default ContentGrid;
