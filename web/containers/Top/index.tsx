import * as React from 'react';
import MainVisual from './MainVisual';
import { Section } from '../../components';
import Footer from './Footer';

const Top = () => (
  <>
    <MainVisual />
    <Section title="NEWS">ここにニュースを入れます</Section>
    <Section title="CONTENTS">
      ここにSessionとかコンテンツ内容が入ります
    </Section>
    <Section title="ACCESS">ここに地図とかが入ります</Section>
    <Footer />
  </>
);

export default Top;