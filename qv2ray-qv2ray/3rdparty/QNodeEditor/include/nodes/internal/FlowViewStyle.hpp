#pragma once
#include "Export.hpp"
#include "Style.hpp"

#include <QtGui/QColor>
namespace QtNodes
{
    class NODE_EDITOR_PUBLIC FlowViewStyle : public Style
    {
      public:
        FlowViewStyle();
        FlowViewStyle(QString jsonText);

      public:
        static void setStyle(QString jsonText);
        static void reset();

      private:
        void resetStyle() override;
        void loadJsonText(QString jsonText) override;
        void loadJsonFile(QString fileName) override;
        void loadJsonFromByteArray(QByteArray const &byteArray) override;

      public:
        QColor BackgroundColor;
        QColor FineGridColor;
        QColor CoarseGridColor;
    };
} // namespace QtNodes
