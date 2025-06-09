import ChannelChat from "../../../../components/ChannelChat";

export default function Page({ params }) {
    const { serverId, channelId } = params;
    return <ChannelChat serverId={serverId} channelId={channelId}/>;
}
